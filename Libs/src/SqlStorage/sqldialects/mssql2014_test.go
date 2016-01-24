package SqlDialects

import (
	"testing"
	"time"
)

func TestConvertSomethingIntoMssql2014SqlScriptString(t *testing.T) {

	expected := SqlScriptString("NULL")
	actual, err := convertSomethingIntoMssql2014SqlScriptString(nil)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(nil) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(nil) returned `%#v` while expected `%#v`", actual, expected)
	}

	str := "Test String havin' lots of characters  !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~ ¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘęĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲųŴŵŶŷŸŹźŻżŽžſƀƁƂƃƄƅƆƇƈƉƊƋƌƍƎƏƐƑƒƓƔƕƖƗƘƙƚƛƜƝƞƟƠơƢƣƤƥƦƧƨƩƪƫƬƭƮƯưƱƲƳƴƵƶƷƸƹƺƻƼƽƾƿǀǁǂǃǄǅǆǇǈǉǊǋǌǍǎǏǐǑǒǓǔǕǖǗǘǙǚǛǜǝǞǟǠǡǢǣǤǥǦǧǨǩǪǫǬǭǮǯǰǱǲǳǴǵǶǷǸǹǺǻǼǽǾǿȀȁȂȃȄȅȆȇȈȉȊȋȌȍȎȏȐȑȒȓȔȕȖȗȘșȚțȜȝȞȟȠȡȢȣȤȥȦȧȨȩȪȫȬȭȮȯȰȱȲȳȴȵȶȷȸȹȺȻȼȽȾȿɀɁɂɃɄɅɆɇɈɉɊɋɌɍɎɏɐɑɒɓɔɕɖɗɘəɚɛɜɝɞɟɠɡɢɣɤɥɦɧɨɩɪɫɬɭɮɯɰɱɲɳɴɵɶɷɸɹɺɻɼɽɾɿʀʁʂʃʄʅʆʇʈʉʊʋʌʍʎʏʐʑʒʓʔʕʖʗʘʙʚʛʜʝʞʟʠʡʢʣʤʥʦʧʨʩʪʫʬʭʮʯʰʱʲʳʴʵʶʷʸʹʺʻʼʽʾʿˀˁ˂˃˄˅ˆˇˈˉˊˋˌˍˎˏːˑ˒˓˔˕˖˗˘˙˚˛˜˝˞˟ˠˡˢˣˤ˥˦˧˨˩˪˫ˬ˭ˮ˯˰˱˲˳˴˵˶˷˸˹˺˻˼˽˾˿̴̵̶̷̸̡̢̧̨̛̖̗̘̙̜̝̞̟̠̣̤̥̦̩̪̫̬̭̮̯̰̱̲̳̹̺̻̼͇͈͉͍͎̀́̂̃̄̅̆̇̈̉̊̋̌̍̎̏̐̑̒̓̔̽̾̿̀́͂̓̈́͆͊͋͌̕̚ͅ͏͓͔͕͖͙͚͐͑͒͗͛ͣͤͥͦͧͨͩͪͫͬͭͮͯ͘͜͟͢͝͞͠͡ͰͱͲͳʹ͵Ͷͷ͸͹ͺͻͼͽ;Ϳ΀΁΂΃΄΅Ά·ΈΉΊ΋Ό΍ΎΏΐΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡ΢ΣΤΥΦΧΨΩΪΫάέήίΰαβγδεζηθικλμνξοπρςστυφχψωϊϋόύώϏϐϑϒϓϔϕϖϗϘϙϚϛϜϝϞϟϠϡϢϣϤϥϦϧϨϩϪϫϬϭϮϯϰϱϲϳϴϵ϶ϷϸϹϺϻϼϽϾϿЀ"
	expected = SqlScriptString("N'Test String havin'' lots of characters  !\"#$%&''()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~ ¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘęĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲųŴŵŶŷŸŹźŻżŽžſƀƁƂƃƄƅƆƇƈƉƊƋƌƍƎƏƐƑƒƓƔƕƖƗƘƙƚƛƜƝƞƟƠơƢƣƤƥƦƧƨƩƪƫƬƭƮƯưƱƲƳƴƵƶƷƸƹƺƻƼƽƾƿǀǁǂǃǄǅǆǇǈǉǊǋǌǍǎǏǐǑǒǓǔǕǖǗǘǙǚǛǜǝǞǟǠǡǢǣǤǥǦǧǨǩǪǫǬǭǮǯǰǱǲǳǴǵǶǷǸǹǺǻǼǽǾǿȀȁȂȃȄȅȆȇȈȉȊȋȌȍȎȏȐȑȒȓȔȕȖȗȘșȚțȜȝȞȟȠȡȢȣȤȥȦȧȨȩȪȫȬȭȮȯȰȱȲȳȴȵȶȷȸȹȺȻȼȽȾȿɀɁɂɃɄɅɆɇɈɉɊɋɌɍɎɏɐɑɒɓɔɕɖɗɘəɚɛɜɝɞɟɠɡɢɣɤɥɦɧɨɩɪɫɬɭɮɯɰɱɲɳɴɵɶɷɸɹɺɻɼɽɾɿʀʁʂʃʄʅʆʇʈʉʊʋʌʍʎʏʐʑʒʓʔʕʖʗʘʙʚʛʜʝʞʟʠʡʢʣʤʥʦʧʨʩʪʫʬʭʮʯʰʱʲʳʴʵʶʷʸʹʺʻʼʽʾʿˀˁ˂˃˄˅ˆˇˈˉˊˋˌˍˎˏːˑ˒˓˔˕˖˗˘˙˚˛˜˝˞˟ˠˡˢˣˤ˥˦˧˨˩˪˫ˬ˭ˮ˯˰˱˲˳˴˵˶˷˸˹˺˻˼˽˾˿̴̵̶̷̸̡̢̧̨̛̖̗̘̙̜̝̞̟̠̣̤̥̦̩̪̫̬̭̮̯̰̱̲̳̹̺̻̼͇͈͉͍͎̀́̂̃̄̅̆̇̈̉̊̋̌̍̎̏̐̑̒̓̔̽̾̿̀́͂̓̈́͆͊͋͌̕̚ͅ͏͓͔͕͖͙͚͐͑͒͗͛ͣͤͥͦͧͨͩͪͫͬͭͮͯ͘͜͟͢͝͞͠͡ͰͱͲͳʹ͵Ͷͷ͸͹ͺͻͼͽ;Ϳ΀΁΂΃΄΅Ά·ΈΉΊ΋Ό΍ΎΏΐΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡ΢ΣΤΥΦΧΨΩΪΫάέήίΰαβγδεζηθικλμνξοπρςστυφχψωϊϋόύώϏϐϑϒϓϔϕϖϗϘϙϚϛϜϝϞϟϠϡϢϣϤϥϦϧϨϩϪϫϬϭϮϯϰϱϲϳϴϵ϶ϷϸϹϺϻϼϽϾϿЀ'")
	actual, err = convertSomethingIntoMssql2014SqlScriptString(str)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(str) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(str) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&str)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&str) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&str) returned `%#v` while expected `%#v`", actual, expected)
	}

	i := 2147483647
	expected = "2147483647"
	actual, err = convertSomethingIntoMssql2014SqlScriptString(i)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(i) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(i) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&i)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&i) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&i) returned `%#v` while expected `%#v`", actual, expected)
	}

	bi := int64(9223372036854775807)
	expected = "9223372036854775807"
	actual, err = convertSomethingIntoMssql2014SqlScriptString(bi)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(bi) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(bi) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&bi)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&bi) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&bi) returned `%#v` while expected `%#v`", actual, expected)
	}

	f32 := float32(2234.1235)
	expected = "2234.1235"
	actual, err = convertSomethingIntoMssql2014SqlScriptString(f32)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(f32) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(f32) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&f32)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&f32) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&f32) returned `%#v` while expected `%#v`", actual, expected)
	}

	f64 := float64(123456789.12345678)
	expected = "123456789.12345678"
	actual, err = convertSomethingIntoMssql2014SqlScriptString(f64)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(f64) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(f64) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&f64)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&f64) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&f64) returned `%#v` while expected `%#v`", actual, expected)
	}

	tm, _ := time.Parse(MSSQL2014_TIMEPARSE_TEMPLATE, "20150908 23:59:59.8888888 -02:30")
	expected = "TRY_CAST('20150908 23:59:59.8888888 -02:30' AS DATETIMEOFFSET)"
	actual, err = convertSomethingIntoMssql2014SqlScriptString(tm)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(tm) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(tm) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&tm)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&tm) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&tm) returned `%#v` while expected `%#v`", actual, expected)
	}

	b := true
	expected = "1"
	actual, err = convertSomethingIntoMssql2014SqlScriptString(b)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(b) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(b) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&b)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&b) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&b) returned `%#v` while expected `%#v`", actual, expected)
	}

	b = false
	expected = "0"
	actual, err = convertSomethingIntoMssql2014SqlScriptString(b)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(b) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(b) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&b)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&b) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&b) returned `%#v` while expected `%#v`", actual, expected)
	}

	sbt := []byte("The TestString")
	expected = "0x5468652054657374537472696e67"
	actual, err = convertSomethingIntoMssql2014SqlScriptString(sbt)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(sbt) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(sbt) returned `%#v` while expected `%#v`", actual, expected)
	}
	actual, err = convertSomethingIntoMssql2014SqlScriptString(&sbt)
	if err != nil {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&sbt) returned unexpected error: %#v", err)
	}
	if actual != expected {
		t.Errorf("convertSomethingIntoMssql2014SqlScriptString(&sbt) returned `%#v` while expected `%#v`", actual, expected)
	}

}
