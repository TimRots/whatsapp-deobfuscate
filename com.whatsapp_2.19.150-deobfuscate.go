/*
29/05/2019 - Tim Rots
com.whatsapp-2.19.150

While auditing the com.whatsapp package I stumbled upon the
instance method which handles deobfuscation of the app secrets
which are referenced at multiple locations trough out the app.


package p085d.p241f.p305X;

public class C4218b {
    public static final String f14658a = C4218b.m8240a("zffb(==q~{w|fa!<u}}u~w<q}=uw|w`sfwM \"&");
    public static final String f14659b = C4218b.m8240a("zffba(==d<ezsfasbb<|wf");
    public static final String f14660c = C4218b.m8240a("zffba(==d<ezsfasbb<|wf=d =wj{af");
    public static final String f14678u = C4218b.m8240a("esb`{dsqk?fw|}`?%%%$ $**'<ga?wsaf?#<w~p<ssh}|sea<q}");
    ..etc..
    public static String m8240a(String str) {
        StringBuilder stringBuilder = new StringBuilder();
        for (int i = 0; i < str.length(); i++) {
            stringBuilder.append((char) (str.charAt(i) ^ 18));
        }
        return stringBuilder.toString();
    }

I have ported the class and m8240a method to Go for educational and research purposes.

Build and run as follow:

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

I am not sure if this is a proper way of storing app secrets..
*/

package main

import "fmt"

func whatsappDeobfuscate(secret string) {
	var result string
	for pos, char := range secret {
		fmt.Printf("[%d] %#U \t", pos, char^18)
		result += string(char ^ 18)
	}
	fmt.Printf("\n\nDeobfuscated: %s\n\n", result)
}

func main() {
	secrets := []string{
	    "zffb(==q~{w|fa!<u}}u~w<q}=uw|w`sfwM \"&",
	    "zffba(==d<ezsfasbb<|wf",
	    "zffba(==d<ezsfasbb<|wf=d =wj{af",
	    "zffba(==d<ezsfasbb<|wf=d =q}vw",
	    "zffba(==d<ezsfasbb<|wf=d =`wu{afw`",
	    "zffba(==d<ezsfasbb<|wf=d =awqg`{fk",
	    "zffba(==eee<ezsfasbb<q}=afsfga<bzb-d/ ",
	    "zffba(==sb{<t}g`acgs`w<q}=d =dw|gwa=aws`qz",
	    "BTGXK T^WJK^PQJZW@UTH[CDE'\"#[DJQJ_JAZJ\\VEV[JGCSF",
	    "JCXS\"ZE K#Z[KB\\ VP_E@!VBBVFZE W#D!!B^STD@XTW_XA#",
	    " \"#&\"$\"#",
	    "zffba(==u`sbz<tsqwp}}y<q}=d <!=",
	    "#$\"+& %*\"'+''\" &nt#vw$tqvqp##p #'ws%s v!qv\"$ wqtt",
	    "zffba(==eee<p{|usb{a<q}=sb{=d$={suwa=aws`qz",
	    "Af`{qf",
	    "V&#V*QV+*T\"\"P \"&W+*\"\"++*WQT*& %W&T&S%&+ ",
	    "Acgs`w",
	    "Sb|^[[[hdxq@Q?t*ZdW\"~JPkd$KJ*VWfQY$!HCGcEbXgdqABFg+`vfbc\\GZKU!V",
	    "^_&\\uqxT?JMWwvsH$_eBDAZ\"Asq/",
	    "A\u0004\u001d@\u0011\u0018V\u0002T(3{;ES",
	    "esb`{dsqk?fw|}`?%%%$ $**'<ga?wsaf?#<w~p<ssh}|sea<q}",
	    "esb`{dsqk?u{bzk?$+\"*$% +\"<ga?wsaf?#<w~p<ssh}|sea<q}",
	    "\\J H_  C#P![",
	    "u!V!@~zBzg]S",
	    "&x1w8T+9_a7nu#l'<!`Z3ew>",
	    "\f\u000f\u001eªåGvJ\u0011W,jÔÇ",
	    "#&\"*'''# !$",
	}
	for secret := range secrets {
		whatsappDeobfuscate(secrets[secret])
	}
}
