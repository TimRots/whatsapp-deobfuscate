# whatsapp-deobfuscate

29/05/2019 - Tim Rots / com.whatsapp-2.19.150

poc for WhatsApp for android app secret deobfuscation 

```
$ go build com.whatsapp_2.19.150-deobfuscate.go;./com.whatsapp_2.19.150-deobfuscate|grep Deobfuscated
Deobfuscated: http://clients3.google.com/generate_204
Deobfuscated: https://v.whatsapp.net
Deobfuscated: https://v.whatsapp.net/v2/exist
Deobfuscated: https://v.whatsapp.net/v2/code
Deobfuscated: https://v.whatsapp.net/v2/register
Deobfuscated: https://v.whatsapp.net/v2/security
Deobfuscated: https://www.whatsapp.com/status.php?v=2
Deobfuscated: https://api.foursquare.com/v2/venues/search
Deobfuscated: PFUJY2FLEXYLBCXHERGFZIQVW501IVXCXMXSHXNDWDIXUQAT
Deobfuscated: XQJA0HW2Y1HIYPN2DBMWR3DPPDTHW2E1V33PLAFVRJFEMJS1
Deobfuscated: 20140601
Deobfuscated: https://graph.facebook.com/v2.3/
Deobfuscated: 1609427805955024|f1de6fcdcb11b215ea7a2d3cd062ecff
Deobfuscated: https://www.bingapis.com/api/v6/images/search
Deobfuscated: Strict
Deobfuscated: D41D8CD98F00B204E9800998ECF8427E4F4A7492
Deobfuscated: Square
Deobfuscated: ApnLIIIzvjcRC-f8HvE0lXByv6YX8DEtCK63ZQmUqWpJuvcSPTu9rdtpqNUHYG3D
Deobfuscated: LM4NgcjF-X_EedaZ6MwPVSH0Sac=
Deobfuscated: SR
Deobfuscated: waprivacy-tenor-777626885.us-east-1.elb.amazonaws.com
Deobfuscated: waprivacy-giphy-690867290.us-east-1.elb.amazonaws.com
Deobfuscated: NX2ZM22Q1B3I
Deobfuscated: g3Dm3RlhPhuOA
Deobfuscated: 4j#e*F9+Ms%|g1~5.3rH!we,
```
