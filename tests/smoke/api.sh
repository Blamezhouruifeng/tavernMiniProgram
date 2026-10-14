#!/bin/bash
# 酒馆小程序后端冒烟测试（先付制全链路）
# 用法：bash tests/smoke/api.sh [baseUrl]   默认 http://127.0.0.1:8000
set -e
BASE=${1:-http://127.0.0.1:8000}
B=$BASE/api/tavern
DIE() { echo "❌ $1"; exit 1; }
OK() { echo "✅ $1"; }

login() {
  curl -s -X POST $B/member/login -H 'Content-Type: application/json' -d "{\"code\":\"$1\",\"nickname\":\"冒烟测试\"}" \
    | python3 -c 'import sys,json;print(json.load(sys.stdin)["data"]["token"])'
}

TOKEN=$(login "smoke-$RANDOM")
AUTH="Authorization: Bearer $TOKEN"
[ -n "$TOKEN" ] || DIE "会员登录失败"
OK "会员登录"

MENU=$(curl -s -H "$AUTH" "$B/menu/list")
echo "$MENU" | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"]["list"];sys.exit(0 if d else 1)' || DIE "菜单为空"
OK "菜单加载"

SID=$(curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' -d '{"tableNo":"A01"}' $B/session/open \
  | python3 -c 'import sys,json;print(json.load(sys.stdin)["data"]["id"])')
[ -n "$SID" ] || DIE "开台失败"
OK "扫码开台 sessionId=$SID"

curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' \
  -d "{\"sessionId\":$SID,\"goodsId\":201,\"skuId\":2012,\"qty\":1}" $B/cart/add > /dev/null
OK "加入购物车"

# 只取「自己的篮子」— 同一桌可能有其他会员的购物车行（全桌共享可见）
MY_ID=$(curl -s -H "$AUTH" "$B/member/profile" | python3 -c 'import sys,json;print(json.load(sys.stdin)["data"]["id"])')
CART=$(curl -s -H "$AUTH" "$B/session/snapshot?sessionId=$SID" \
  | python3 -c "import sys,json;d=json.load(sys.stdin)['data']['cart'];print(json.dumps([c['id'] for c in d if c['memberId']==$MY_ID]))")
COUPON=$(curl -s -H "$AUTH" "$B/coupon/list?status=1" | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"]["list"];print(d[0]["id"] if d else 0)')

RESP=$(curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' \
  -d "{\"sessionId\":$SID,\"cartIds\":$CART,\"couponId\":$COUPON}" $B/order/submit)
ON=$(echo "$RESP" | python3 -c 'import sys,json;print(json.load(sys.stdin)["data"]["orderNo"])')
[ -n "$ON" ] || DIE "提交订单失败"
echo "$RESP" | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"];print("   订单",d["orderNo"],"总额",d["totalAmount"],"优惠",d["discountAmount"],"应付",d["payAmount"])'
OK "提交订单（先付制待支付）"

PAID=$(curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' -d "{\"orderNo\":\"$ON\"}" $B/order/pay)
echo "$PAID" | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"];sys.exit(0 if d["payStatus"]==1 and d["makingStatus"]==1 else 1)' \
  || DIE "支付后未进入制作"
OK "支付成功即进入制作（无人工接单）"

curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' -d '{"isStaff":true,"role":"waiter"}' $B/member/switchStaff > /dev/null
curl -s -H "$AUTH" "$B/staff/tickets" | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"]["list"];sys.exit(0 if d else 1)' || DIE "出票区为空"
OK "员工出票区可见已支付订单"

curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' -d "{\"orderNo\":\"$ON\"}" $B/staff/serve \
  | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"];sys.exit(0 if d["makingStatus"]>=2 else 1)' || DIE "行级上齐失败"
OK "行级上齐"

ITEM=$(curl -s -H "$AUTH" "$B/order/detail?orderNo=$ON" | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"];print([i["id"] for i in d["items"] if i["isWine"]][0])')
OID=$(curl -s -H "$AUTH" "$B/order/detail?orderNo=$ON" | python3 -c 'import sys,json;print(json.load(sys.stdin)["data"]["id"])')
curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' -d "{\"orderId\":$OID,\"itemIds\":[$ITEM]}" $B/bottle/depositFromOrder \
  | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"];sys.exit(0 if d["remainRatio"]==100 else 1)' || DIE "退酒入存失败"
OK "退酒入存"

DEP=$(curl -s -H "$AUTH" "$B/bottle/list" | python3 -c 'import sys,json;print(json.load(sys.stdin)["data"]["list"][0]["id"])')
CODE=$(curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' -d "{\"depositIds\":[$DEP]}" $B/bottle/takeCode \
  | python3 -c 'import sys,json;print(json.load(sys.stdin)["data"]["code"])')
curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' -d "{\"code\":\"$CODE\",\"depositId\":$DEP,\"takeRatio\":30}" $B/staff/bottle/verify \
  | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"];sys.exit(0 if d["deposit"]["remainRatio"]==70 else 1)' || DIE "存酒核销失败（剩余比例不为 70）"
OK "存酒核销（剩余 70%）"

curl -s -X POST -H "$AUTH" -H 'Content-Type: application/json' -d "{\"orderNo\":\"$ON\",\"reason\":\"冒烟测试自助退\"}" $B/order/refund/apply \
  | python3 -c 'import sys,json;d=json.load(sys.stdin)["data"];sys.exit(0 if d["refundStatus"] in (1,4) else 1)' || DIE "退款申请失败"
OK "退款申请"

echo ""
echo "🎉 冒烟测试全部通过"
