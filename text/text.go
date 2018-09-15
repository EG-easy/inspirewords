package text

import (
	"time"
)

func ChooseTweet()string{
	t := time.Now().Day()
	tweetlist :=  [31]string{
		"【1日】問題は未来だ。だから私は、過去を振り返らない\n" +
			"by ビル・ゲイツ（マイクロソフト創業者）",
		"【2日】（僕は毎日のようにこう自問している）「今僕は自分にできる一番大切なことをやっているのだろうか」と\n" +
			"by マーク・ザッカーバーグ（フェイスブック創業者）",
		"【3日】私たちは世界に役立つことをしている。それが一番大事\n" +
			"byイーロン・マスク（スペースX、テスラ・モーターズCEO）",
		"【4日】優れたリーダーは、どこでノーというべきかを知っている\n" +
			"by ショーン・パーカー（ナップスター創業者、フェイスブック初代CEO）",
		"【5日】目標に向かう時は、ちょっとまぬけでなくっちゃいけないのさ\n" +
			"by ラリー・ペイジ（グーグルの創業者）",
		"【6日】エクセレントな人間になるためにかかる時間は1分だ\n" +
			"by トーマス・ワトソン（IBM初代社長）",
		"【7日】気持ちを燃えさせよ\n" +
			"by ジャック・ウェルチ（GE元CEO）",
		"【8日】自信を持つことと、うぬぼれることには、微妙な違いがある\n" +
			"by ジョン・スカリー（アップル元CEO）",
		"【9日】金を手にして得意がるのは愚か者だけだ\n" +
			"by ジョン・ロックフェラー（スタンダード石油創業者）",
		"【10日】波が到達するのを待ってはいられない。漕ぎ始めなければ\n" +
			"by ピーター・ティール（ペイパル創業者）",
		"【11日】理屈や常識に反するとしても、決断しなければならない時がある\n" +
			"by ハワード・シュルツ（スターバックスCEO）",
		"【12日】今解決策が見当たらないからといって、それが存在しないわけじゃない\n" +
			"by エリック・シュミット（グーグル元CEO）",
		"【13日】最初の一人は必ず大当たりだ\n" +
			"by ゴートン・ムーア（インテル創業者）",
		"【14日】一日23時間働く覚悟がなかったらやめた方がいい\n" +
			"by フィル・ナイト（ナイキ創業者）",
		"【15日】51対49じゃない。98対2で勝つ\n" +
			"by トラビス・カラニック（ウーバー創業者）",
		"【16日】逆境の中で咲く花は、どの花よりも貴重で美しい\n" +
			"by ウォルト・ディズニー（ディズニー創業者）",
		"【17日】世界を急に変えることはできなくても、少しずつ良くしていくことはできる\n" +
			"by スティーブ・ウォズニアック（アップル創業者）",
		"【18日】やり遂げろ、この世界で継続ほど価値のあるものはない\n" +
			"by レイ・クロック（マクドナルドコーポレーション創業者）",
		"【19日】贅沢にとらわれることは、終わりの始まり\n" +
			"by サム・ウォルトン（ウォルマート創業者）",
		"【20日】早く失敗したほうがいい。そして過ちを小規模に留めなさい\n" +
			"by ニクラス・ゼンストローム（スカイプ創業者）",
		"【21日】誰かがすごいことをやった時は、なぜ俺じゃないのか、と腹が立つ\n" +
			"by ラリー・エリソン（オラクル創業者）",
		"【22日】名声を打ち立てるには一生かかるが、台なしにするには5分とかからない\n" +
			"by ウォーレン・バフェット（バークシャー・ハサウェイCEO、世界一の投資家）",
		"【23日】仕事がうまくいっても10億分の1秒だけ祝って先に進むのだ\n" +
			"by マイケル・デル（デル創業者）",
		"【24日】私に一番大切な仕事は、仕事を楽しくすることだ\n" +
			"by アンドリュー・グローブ（インテル創立メンバー）",
		"【25日】利益だけを考えてはいけない。世の中を変えることに重点を置くのです。\n" +
			"by ジェリー・ヤン（ヤフー創業者）",
		"【26日】賢くあるより優しくあるほうが難しい\n" +
			"by ジョフ・ベゾス（アマゾン創業者）",
		"【27日】誤っているものはたいていアンフェアですし、正しいものはフェアです\n" +
			"by ジミー・ウェールズ（ウィキペディア創業者）",
		"【28日】スタートは選べない。しかし、「どれだけ努力するか」は自由に制御できる\n" +
			"by　マイケル・ブルームバーグ（ブルームバーグ創業者、元ニューヨーク市長）",
		"【29日】主役は君たちだ。経営者が映画をつくるんじゃない\n" +
			"by ジョン・ラセター（ピクサーメンバー、ディズニーアニメーション部門責任者）",
		"【30日】過去や未来に捉われずに今を生きる\n" +
			"by ジャック・ドーシー（ツイッター創業者）",
		"【31日】ハングリーになれ！愚か者になれ！\n" +
			"by スティーブ・ジョブズ（アップル創業者）",
	}
	return  tweetlist[t]
}
