package whitelist

type Protocol struct {
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	Twitter   string   `json:"twitter,omitempty"`
	Logo      string   `json:"logo,omitempty"`
	Contracts []string `json:"contracts"`
}

var (
	Contracts = `[
  {
    "name": "NFTs2Me",
    "slug": "nfts-2-me",
    "twitter": "@NFTs2Me",
    "logo": "nfts2me.png",
    "contracts": [
      "0x00000000001594C61dD8a6804da9AB58eD2483ce"
    ]
  },
  {
    "name": "Magpie Protocol",
    "slug": "magpie-protocol",
    "contracts": [
      "0x956df8424b556f0076e8abf5481605f5a791cc7f"
    ]
  },
  {
    "name": "Aspecta",
    "slug": "aspecta",
    "contracts": [
      "0x3636814Ab8d82854949d5794c428B8329115755D"
    ]
  },
  {
    "name": "Rubic",
    "slug": "Rubic",
    "logo": "rubic.png",
    "contracts": [
         "0xAa4472EC72cF4771dfD38467f161F6DF6cA3FB1a",
	 "0xAf14797CcF963B1e3d028a9d51853acE16aedBA1",
         "0x3335733c454805df6a77f825f266e136FB4a3333",
	 "0xf747DA269c7a65d123d4851515927b082e23aBd2",
         "0xce0f6610fc406f689C854Ac517eA15D0205c11E3",
	 "0x40785ea9e20B7F93ab0d9DCe73B0775d52901e66",
         "0xc4A04649E0e2bCC5859F8aAa7c7A79a1a0cC388d",
         "0xC9fa6Ca19d49ff5a677181369eB8a58688E74924",
         "0x6535B3E3802941B91405306CDF6c89A45875Fc98",
         "0xb656E75248Ee97e14Dccb5EC559Da32a5b70c8eF",
         "0x1C7cc271Ad9DBe625a07C653aE61b757fdFCbe91"
    ]
  },
  {
    "name": "Bunnyfi",
    "slug": "bunnyfi",
    "twitter": "@bunnyfilabs",
    "logo": "bunnyfi.svg",
    "contracts": [
      "0x0000000000D310F8802cC91F198d14bC2303230B"
    ]
  },
  {
    "name": "Kodo Exchange",
    "slug": "kodo-exchange",
    "twitter": "@kodohq",
    "logo": "kodo.png",
    "contracts": [
      "0x7E034Ef620D2fb403e8bB6a1130670110287A7a1",
      "0x7e91F29F8a213c8311712A8FC8c61219fb9477CB",
      "0x6c4A102B7aafFA9a8C9440c08A5c09deECAFB324",
      "0x535E02960574d8155596a73c7Ad66e87e37Eb6Bc",
      "0x7149E14784f9d88B5497a9bf135c643151379F95",
      "0x3a9E14D73AD40E70baFaFfefE8893Eb318Fc2312",
      "0xd04d75E1CDe512b195E70C6c18Cf7Ec4b2B12f41",
      "0xbf6fabcc707aC239Be2D7818797745F678A411ad",
      "0x1A805BBcE7F87365daC956cFD8d078ef827E73d1",
      "0x0e16aA850AF7956B476Ad6056ead67A32f099504",
      "0x8ba3C594Dc3796c171a1B7F0e143577abE03300F",
      "0x46E9cef07e01ab5a73E9B10cfb423E9319cD68c9"
    ]
  },
  {
    "name": "Micro3",
    "slug": "micro-3",
    "twitter": "@micro3io",
    "logo": "micro3.jpg",
    "contracts": [
      "0xdCe78c98A0a5aBE3f0A342b32BfE43D203794DeE",
      "0x7b0A5F84181425C58a61659BE9952a04a9a9A833",
      "0xfE6AF84f2fA6f2C24c7afB05337f7c2527dD2709",
      "0x1e5a9ad59434c59d7a5d19bc3e5dad1d94255a3f",
      "0xaa52404f84b17f7a3f9b15e4981c5ff6e8d33786",
      "0x6c18486684e74987800665222029424eea56e6e8",
      "0x330678c37aacf710fd97b6e75157fc39d82cc610",
      "0x689f6ae83d9c84eb278318f4b9169b2d78f703ae",
      "0x9b348c3ba23605880a39a7141c26675cea184688",
      "0xee3f99089da945d04cba9b1085dbd04e49fa40a4",
      "0x0000C019d60b628F9Ba553092CdA375191319c5e",
      "0xFa8230a55cBFF23299000361bce4160635cE756b",
      "0x2f5c25874302558299e955d29619c6F4b9A16610",
      "0x64C310638a405D15A5C6C4590ac4A1f20d8A7546",
      "0x9f5247723cc31d19ce87943b075f46a8e583afdd",
      "0x81e3d94da4a941410f028685f755679187b28dae",
      "0x4a12dec6d7eac74a36f78f844afa04bfa880d96a",
      "0xa96dede9dd8a821a862b515b317a309cad60e670",
      "0xabd7c118d8d77f13880339de3d237a9acde39b76",
      "0xd82cc6a9a602732faf5c803e017d9c66bfc43b71",
      "0x678e2f54d5a13881faa3cb899311e797c386aa1e"
    ]
  },
  {
    "name": "Comet Protocol",
    "slug": "comet-protocol",
    "twitter": "@Comet_Protocol",
    "logo": "comet.svg",
    "contracts": [
      "0xB50Ac92D6d8748AC42721c25A3e2C84637385A6b",
      "0x0fbCf4a62036E96C4F6770B38a9B536Aa14d1846"
    ]
  },
  {
    "name": "Stargate Finance",
    "slug": "stargate",
    "twitter": "@stargatefinance",
    "logo": "stargate.jpg",
    "contracts": [
      "0x0dB9afb4C33be43a0a0e396Fd1383B4ea97aB10a",
      "0xCd4302D950e7e6606b6910Cd232758b5ad423311",
      "0x711b5aAFd4d0A5b7B863Ca434A2678D086830d8E",
      "0x77C71633C34C3784ede189d74223122422492a0f",
      "0x1C10CC06DC6D35970d1D53B2A23c76ef370d4135",
      "0x45d417612e177672958dC0537C45a8f8d754Ac2E",
      "0x9c2dc7377717603eB92b2655c5f2E7997a4945BD",
      "0x19e26B0638bf63aa9fa4d14c6baF8D52eBE86C5C"
    ]
  },
  {
    "name": "Crack & Stack",
    "slug": "crack-stack",
    "twitter": "@crackandstack",
    "logo": "crackandstack.svg",
    "contracts": [
      "0x7ddB8A975778a434dE03dd666F11Ce962DCdD290",
      "0x2c301eBfB0bb42Af519377578099b63E921515B7",
      "0x6C8865042792B5E919fC95bF771ccaDF6F0cfA22",
      "0xD31A4be996b7E1cc20974181127E6fCA15649913",
      "0xA9EC1fEEE212851c829B028F094156CD04A3a547",
      "0x1ACa21A2a2a070d3536a69733c7044feDEB88f5A",
      "0xb64C1461453DAdD104A583dCCeef30ce296fde20",
      "0xD8F7cd7d919c5266777FB83542F956dD30E80187",
      "0x009C32F03d6eEa4F6DA9DD3f8EC7Dc85824Ae0e6",
      "0xF8F1B21615BDbEA8D142cfaf4828EA0236Cab115",
      "0x12689b6ddE632E69fBAA70d066f86aC9fDd33dd1"
    ]
  },
  {
    "name": "HenjinDEX",
    "slug": "henjin",
    "twitter": "@HenjinDex",
    "logo": "henjin_dex.png",
    "contracts": [
      "0x07Bc9a408B385C7Aa8De2783795759512fE24356",
      "0xcFf128C67bCDc5a7c7D3F24c638e59AA0d4e112b",
      "0xaF0C5CBbCEfB685BF3200684d2aE19B8eA6186ca"
    ]
  },
  {
    "name": "Oku Trade",
    "slug": "oku",
    "twitter": "@okutrade",
    "logo": "oku.jpg",
    "contracts": [
      "0x75FC67473A91335B5b8F8821277262a13B38c9b3",
      "0x346239972d1fa486FC4a521031BC81bFB7D6e8a4",
      "0xB3309C48F8407651D918ca3Da4C45DE40109E641",
      "0xE3dbcD53f4Ce1b06Ab200f4912BD35672e68f1FA",
      "0x454050C4c9190390981Ac4b8d5AFcd7aC65eEffa",
      "0x38EB9e62ABe4d3F70C0e161971F29593b8aE29FF",
      "0x743E03cceB4af2efA3CC76838f6E8B50B63F184c",
      "0x8B3c541c30f9b29560f56B9E44b59718916B69EF",
      "0x6Aa54a43d7eEF5b239a18eed3Af4877f46522BCA",
      "0xaa52bB8110fE38D0d2d2AF0B85C3A3eE622CA455",
      "0x807F4E281B7A3B324825C64ca53c69F0b418dE40",
      "0xdD489C75be1039ec7d843A6AC2Fd658350B067Cf",
      "0x1b35fbA9357fD9bda7ed0429C8BbAbe1e8CC88fc",
      "0xcb2436774C3e191c85056d248EF4260ce5f27A9D",
      "0x447B8E40B0CdA8e55F405C86bC635D02d0540aB8"
    ]
  },
  {
    "name": "ZNS Connect",
    "slug": "zns",
    "twitter": "@ZNSConnect",
    "logo": "zns.svg",
    "contracts": [
      "0xFb2Cd41a8aeC89EFBb19575C6c48d872cE97A0A5"
    ]
  },
  {
    "name": "RubyScore",
    "slug": "ruby-score",
    "twitter": "@rubyscore_io",
    "logo": "rubyscore.svg",
    "contracts": [
      "0xDC3D8318Fbaec2de49281843f5bba22e78338146",
      "0x4D1E2145082d0AB0fDa4a973dC4887C7295e21aB",
      "0x3d52d95D58fCb53814ea37d580601D2AF2B4CC98"
    ]
  },
  {
    "name": "Orbiter Finance",
    "slug": "orbiter",
    "twitter": "@Orbiter_Finance",
    "logo": "orbiter.png",
    "contracts": [
      "0x80c67432656d59144ceff962e8faf8926599bcf8",
      "0xe4edb277e41dc89ab076a1f049f4a3efa700bce8",
      "0xee73323912a4e3772b74ed0ca1595a152b0ef282",
      "0x41d3d33156ae7c62c094aae2995003ae63f587b3",
      "0xd7aa9ba6caac7b0436c91396f22ca5a7f31664fc"
    ]
  },
  {
    "name": "Hana Finance",
    "slug": "hana-finance",
    "twitter": "@Hana_Finance",
    "logo": "hana_finance.png",
    "contracts": [
      "0xB9eD09af341a59c05c8AaE584172e8dCc1E828b6",
      "0x4aB85Bf9EA548410023b25a13031E91B4c4f3b91",
      "0x5C9bC967E338F48535c3DF7f80F2DB0A366D36b2",
      "0xacd2E13C933aE1EF97698f00D14117BB70C77Ef1",
      "0x0247606c3d3f62213bbc9d7373318369e6860eb1",
      "0xf1777EAD4098F574c68E59905588f3C9875251ed"
    ]
  },
  {
    "name": "rhino.fi",
    "slug": "rhino-fi",
    "twitter": "@RhinoFi",
    "logo": "rhinofi.png",
    "contracts": [
      "0x1Df2De291F909baA50C1456C87C71Edf9Fb199D5"
    ]
  },
  {
    "name": "DTX.TRADE",
    "slug": "dtx",
    "twitter": "@0xdtx",
    "logo": "dtx-trade.png",
    "contracts": [
      "0xa814273254C1F73fF79f4D5b5d41279dCbb83f9E",
      "0xCFC8D41f93FCF867f5C2D8435DF84c71672728f6",
      "0x309B732B115f45B41AD3C1F6D16d475974E078F5",
      "0x9702653B16617Dfe5c24c09d11cCBaC8C3Ba1A02",
      "0x061C1D8D66476FF4A6069Db0E780F2EEa26D9bee",
      "0x1aDcbDc6D1e178D12021d33E52Cff85902B77Ce2",
      "0x8f3e1de971FC4D351a8750c8f490E98A83Bc5F8D",
      "0x5E90EC757A38f328f1153A0a2ABda5659B6c9344",
      "0xB221550692aca84fBfD9F0c38C5F9899CabdB436",
      "0x92350121742457A29E8e3F555fd955D0C581816D",
      "0xD1cA9Bae9716fC61Fa1476BE95148945850149D9",
      "0x3db16347DD51093329c1e5Ae3F3f3F844A6890f1",
      "0xeb7794536443ac7603478eC716765BC58Ab7FCf5",
      "0xaa8F179447E00Acd547b755f8F32af0acb939AB8",
      "0x145A0dB7CC5B1EDeb36749F2ED2be6FBEb275ef4",
      "0xc0ab776604059D10880dbD219758FF7B82997cc0",
      "0x1d22B869C281ed16DE50989b5453Ee5cf48B12b4",
      "0x14730A861586215d8133DF5c30cb77cA989B299d",
      "0x488a634861C943236eb5A6A104b223D9b914bc3b",
      "0x275FF25FAC6901557BF504289Ac82dAEA9Eb59ef",
      "0xB00231B308B01Dbb90f16F966F62d86fBc78c450",
      "0x22A4B5493c25Ee8Deb478021180a7f3A83A59296",
      "0xDf0c1038A89A68ed4Ac43070888d2d606cE4fEb2",
      "0x81a5626041e20Cfd2e7445e33882be14EA2E7D89",
      "0xcd5212318e94792F2D4455Ef3d8417fec9aE1824",
      "0xD3325F3FD92b2BFa54A49eAB7241eC0ce8A7131d",
      "0x0B095e4E19DF76F1f0B9bBCe41F4BdfaE50c55e1",
      "0x8AA69Cd8782cEA61525739ef670d67Ea6250692c",
      "0xBdf1216844aB32e380c262dAA4ca3F8597F38FC6",
      "0xCFC48386D10DD7c6F0d2C318A223Cf8296291A7C",
      "0x9D8406140573239FB7EB35B37E17128CdDB414DD",
      "0x903Af86AbE346F1B2Bc0e928f57CFd1D5D49Ab2E",
      "0xdA6a745740Bbdbe5F588b79FEe57f2e10ad4Da11",
      "0xE3e6818bbC193D454f38772D34FA4cf8C19684d5",
      "0x2EA9051d5a48eA2350b26306f2b959D262cf67e1",
      "0x66c43b1CB7aD2f824F90cF08bA05847281721D82",
      "0x6dfD98E8a0ff078A49AF4a08B201F6f63F8d37D3",
      "0x092fCe9F5143B591c08c0BFc46AA8ad8E5ACd1fc",
      "0x019D7bA2E124799Feef405133FE62371108745d6"
    ]
  },
  {
    "name": "RetroBridge",
    "slug": "retrobridge",
    "twitter": "@Retro_Bridge",
    "logo": "retrobridge.jpg",
    "contracts": [
      "0x009905bf008CcA637185EEaFE8F51BB56dD2ACa7"
    ]
  },
  {
    "name": "0xAstra",
    "slug": "0xastra",
    "twitter": "@0xAstra_xyz",
    "logo": "0xastra.svg",
    "contracts": [
      "0x90CE48ED68C6FCAe6F13b445F1573f003cF1804d",
      "0x34723B92aE9708BA33843120A86035D049dA7dfA",
      "0x445ab115F67b1EA7Fa4B5956DeB03796D6CB8e13",
      "0xd0989635D363Bd85ecCE495fe78dF43A9067867D",
      "0xF3f9c67CDd87De0597d607B5fC206299eB64fffa"
    ]
  },
  {
    "name": "Ritsu Protocol",
    "slug": "ritsu",
    "twitter": "@ritsuprotocol",
    "logo": "ritsu-protocol.png",
    "contracts": [
      "0x7160570BB153Edd0Ea1775EC2b2Ac9b65F1aB61B",
      "0x12AF3Ec993EC5d5bD789b3e989c9E95A2F6c586D",
      "0xDFFee0ad5C833f2A5E610dfe9FD1aD82743eA74e",
      "0x3e846B1520E74728EFf445F1f86D348755F738d9",
      "0x0A78CAB89a069555a18B78537f09fab24c03dECd",
      "0xB4F80b81a82184D754F933b0A6C2Ba8D5495567C",
      "0x7c38E9389B27668280E5aaAc372eBCb2ECc1c5E0",
      "0xE75bfdbBE463A4c562F69B45f3A302faC4BB9E16",
      "0x1c91eEe74061C8d888cA02324029190466E2f4d4",
      "0x424Fab7bfA3E3Dd0e5BB96771fFAa72fe566200e",
      "0x53487d1668aFd3FBAD276BEE52cb88c01F72b724",
      "0x3BEbD0720F857DeF80af8dd44B5970A3749743Bc",
      "0xeF4a016F3E54c4520220adE7a496842ECbF83E09",
      "0x6c7839E0CE8AdA360a865E18a111A462d08DC15a",
      "0x0627A336b351fE74f6aa5b37F9c1DfCE29192BA1",
      "0x608Cb7C3168427091F5994A45Baf12083964B4A3",
      "0xa2A09f15c2ec6aF1b8f9413c148334b231410bd8",
      "0x80a249A3d22c47C49978e32b7D5C39B37D9b28ad",
      "0x37BAc764494c8db4e54BDE72f6965beA9fa0AC2d",
      "0x60df1f3124128B29D3D595D91CD13dE0B84bD997",
      "0xE4CF807E351b56720B17A59094179e7Ed9dD3727",
      "0x40d660504eB163708d8AC8109fc8F2c063ddAE4b",
      "0x52De0c13b02E3f1F197D61a5f6F04e15881002a8",
      "0xef513EBE1089159C9b7d5492F8EEEb0973094436"
    ]
  },
  {
    "name": "TaikoTown: A LooperLands Experience",
    "slug": "taiko-town",
    "twitter": "@LooperLands",
    "logo": "looperlands.svg",
    "contracts": [
      "0xEe01C4b0538849bF1c66bDFB458a7de11B1d7424"
    ]
  },
  {
    "name": "XY Finance",
    "slug": "xy-finance",
    "twitter": "@xyfinance",
    "logo": "xy_finance.svg",
    "contracts": [
      "0x73Ce60416035B8D7019f6399778c14ccf5C9c7A1",
      "0xedC061306A79257f15108200C5B82ACc874C239d"
    ]
  },
  {
    "name": "KiloEx",
    "slug": "kiloex",
    "twitter": "@KiloEex_perp",
    "logo": "kiloex.png",
    "contracts": [
      "0xd6e87aCd56F3dcb028A906D957be5f65Fe0CFB93",
      "0x735D00A9368164B9dcB2e008d5Cd15b367649aD5",
      "0x6B9DDC7326c330d9d599f14b6a5370d20d494761",
      "0xC7436E6a14A60809C7Ec724564345b43b4789dd1",
      "0x0A4F9C402BbAFe2A7c0B70d1f71788713f7Bd837",
      "0x00E32aD2f1E0cf94F8e05F5D46905034d76CE2DC",
      "0x82Fe1D02B7cf2Ce3dDCB77b2446d9b4e27d79598"
    ]
  },
  {
    "name": "Tokemon.fun",
    "slug": "tokemon",
    "contracts": [
      "0x91b551666C8e8d7de29990901d0a7dCEc81127b9",
      "0xbAd4bae9574fCd967D1D0e775E9932eAEd789e17"
    ]
  },
  {
    "name": "Stupid Monkeys",
    "slug": "stupid-monkeys",
    "twitter": "@StupidMonkeysRH",
    "logo": "stupidmonkeys.png",
    "contracts": [
      "0xCA99F9DbF4A13D4de05B41a68041dcE7929cb5e0"
    ]
  },
  {
    "name": "Owlto Finance",
    "slug": "owlto-finance",
    "twitter": "@Owlto_Finance",
    "logo": "owlto.png",
    "contracts": [
      "0x5e809A85Aa182A9921EDD10a4163745bb3e36284"
    ]
  },
  {
    "name": "coNFT",
    "slug": "co-nft",
    "twitter": "@ConftApp",
    "logo": "conft.jpg",
    "contracts": [
      "0x3AC934F275172a7fa0C0dD4545305bd5EF82a6F8",
      "0x9059cA87Ddc891b91e731C57D21809F1A4adC8D9",
      "0x6Ce2CFD7674cf47A851690a11d1DB45c6cCBe17A",
      "0x8b0d4F55c487eE2b2975EE59044A508eAb23f9F9"
    ]
  },
  {
    "name": "Symbiosis Finance",
    "slug": "symbiosis-finance",
    "twitter": "@symbiosis_fi",
    "logo": "symbiosis.svg",
    "contracts": [
      "0xda8057acB94905eb6025120cB2c38415Fd81BfEB",
      "0x5Aa5f7f84eD0E5db0a4a85C3947eA16B53352FD4",
      "0x7057aB3fB2BeE9c18e0cDe4240DE4ff7f159E365",
      "0xa0079829B9F1Edc5DD0DE3eC104f281745C4bD81"
    ]
  },
  {
    "name": "Pheasant Network",
    "slug": "pheasant",
    "twitter": "@PheasantNetwork",
    "logo": "pheasant.png",
    "contracts": [
      "0x0890f8A7b193A3eEE810DE3AdcFAd181b9ce294E",
      "0x84F90083e4aA00B5FD4DAaaEEc75bdF8978EDCD2",
      "0x558F7547A472a6897126e20440453e57AC320794",
      "0xb9ACb5601C091B39960a6c4974b979483132B30A",
      "0x27cb8546F60fD5d7869a223F40b8036a9eBe3a4f",
      "0x77201FC74123Ea148C836418a08Da3322B3201D3",
      "0x9F28AC2c1a2A82db54DFED6B9784a7A950EfEc08",
      "0x9E7FCb2c0b8a5461BCc7078a2E37886f254B060b",
      "0x04bc515cF3cbd26EA7c621a6220B5648F764030c"
    ]
  },
  {
    "name": "Mini Bridge",
    "slug": "minibridge",
    "twitter": "@Chaineye_tools",
    "logo": "minibridge.jpeg",
    "contracts": [
      "0x00000000000007736e2F9aA5630B8c812E1F3fc9"
    ]
  },
  {
    "name": "Izumi Finance",
    "slug": "izumi-iziswap",
    "twitter": "@izumi_Finance",
    "logo": "izumi_finance.png",
    "contracts": [
      "0x8c7d3063579BdB0b90997e18A770eaE32E1eBb08",
      "0x4d4673745AAC664eFB9758fdd571F40d78a87bfe",
      "0x32D02Fc7722E81F6Ac60B87ea8B4b63a52Ad2b55",
      "0xF4efDB5A1E852f78e807fAE7100B1d38351e38c7",
      "0xe96526e92ee57bBD468DA1721987aa988b008768",
      "0xbD6abA1Ef82A4cD6e15CB05e95f433ef48dfb5df",
      "0x2C6Df0fDbCE9D2Ded2B52A117126F2Dc991f770f",
      "0x14323AfbC2b82fE58F0D9c203830EE969B4d1bE2",
      "0x04830cfCED9772b8ACbAF76Cfc7A630Ad82c9148",
      "0x33531bDBFE34fa6Fd5963D0423f7699775AacaaF",
      "0x34bc1b87f60e0a30c0e24FD7Abada70436c71406",
      "0x7a524c7e82874226F0b51aade60A1BE4D430Cf0F"
    ]
  },
  {
    "name": "Brigade",
    "slug": "brigade",
    "twitter": "@tacowax",
    "logo": "brigade.png",
    "contracts": [
      "0x8a93AAE6D94680658012B887BfDd981A17661Ef4",
      "0x409395BC4b50A9BbD45a943A8B0D6236E0F83540",
      "0x20F50518188FB3c9F5adff472E056291C4B98ecE",
      "0x0158A4055428b67e286b2627e91120b49ca1146c",
      "0x73716C57f87fFd4135453aBCe6cf61Bb0E99C410",
      "0x03376f22eF7d08CEE420D07207f85E52638A9fCd",
      "0x72dCB9a28bB8EA172B58130d9fd17A6dBE7A9E41",
      "0xB66A56126fAd14563e62BA2Cda658cb97F7a90dE",
      "0x4f576c055f06EFFdF165C5bA014f0b827D47E27B"
    ]
  },
  {
    "name": "OKX NFT Market",
    "slug": "okx-nft-market",
    "twitter": "@okxweb3",
    "logo": "okx.png",
    "contracts": [
      "0xa7FD99748cE527eAdC0bDAc60cba8a4eF4090f7c",
      "0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC"
    ]
  },
  {
    "name": "World Of Dypians",
    "slug": "world-of-dypians",
    "twitter": "@worldofdypians",
    "logo": "dypians.png",
    "contracts": [
      "0xaf33f679be47733bD3aBb5b0b977B6ba3eD8d01E",
      "0x620655Ee8320bA51cf4cc06bf6a7C14022271764",
      "0xCb2Eb4ba62346751F36bA652010b553759141AEE",
      "0xc24728996C974788027871C4b1902d9ac70dBCDd",
      "0xf394504D9B224923C2b569B719EE99640fd585dA",
      "0x02E857a5F93F531aAC2bAD56113B81155DbcF236",
      "0x810C42A71358dc1e0Ecc32815DadD90c586AfD1c",
      "0xdb2E1287AAC9974AB28a66fABF9bCB34C5f37712",
      "0xE026FB242D9523dc8E8d8833F7309dbdbED59d3d",
      "0x94d266D7f7F0548282439c0c62962B048ad86d1a",
      "0x2b7DcAeAe477B3Bb4b8a3a043611fE06804F6f9F",
      "0x67C620BbA764b2a45DF5369E7A03d6bD7A210D47",
      "0xd1aFA3A14572cBb11d9aAb0E6D9b3F7780689fd0",
      "0x842856460c527d37D4B5fbBDc3a0a2c20DAF3281",
      "0x5baDC9630a34aC5Af3BBa448A23C681dAEDA679B",
      "0x565571650491C471F31eEfF60e382D3022ccB717",
      "0xeE80b4711c316D66bE0A0865838a8A45e71618E7"
    ]
  },
  {
    "name": "Bullishs",
    "slug": "bullishs",
    "logo": "bullishs.png",
    "contracts": [
      "0x340C885AC77E4a1d3E9e8f6EedfA7269CB374cb6"
    ]
  },
  {
    "name": "REVOX Lense",
    "slug": "revox-lense",
    "contracts": [
      "0x2c80179883217370F777e76c067Eea91D8283C5C"
    ]
  },
  {
  "name":"Meridian Lend",
  "slug": "meridian-lend",
  "contracts":[
    "0x8Cf3E0e7aE4eB82237d0931388EA72D5649D76e0",
    "0x9Fe3659aFf8CDaf4cd009ef7a63DaB23b3CA15C7",
    "0x414B589fe4F26EBeD8Cb2561Bd6Cae05a8fcC037",
    "0x4Ea9de9bEC4571De26f3328F984192e7EFABB669",
    "0xadbb38C5a39afBAD4988C1B0a15479cB52d7962A",
    "0x1697A950a67d9040464287b88fCa6cb5FbEC09BA",
    "0x440a8C55e7e7FD4D2ab91883C6E2716ccF167BbF",
    "0x94bA46c26278661953538fb4AB1d7353A37187cd",
    "0x1fD5343aFD78B3bE331Fb3d809c9131c8788DFf9",
    "0xd79417be2cE02fc22F2E520869b1DfC02a677751",
    "0x6958972e00842E851f8950ec0B2161B5cEec89D6",
    "0xef6881676d7fC7bd6E09927f833a9ACAAab0bA07",
    "0x00A332879be62E2CaD5f57DfcB0768BD00d56D2A",
    "0xd5FC79CD3942ba39807A82fDF60bFf36841E6329",
    "0x42F48F98E432a3e6Ac2911470202d442D73C2fF0",
    "0xd5FC79CD3942ba39807A82fDF60bFf36841E6329",
    "0x87a8b64b88C4A1B3E0E2f2278046E8efbAa04CC9",
    "0x93335200cdbb8c567A205606D7319CC321D085d0",
    "0xF0E8545d106D548c8688795dD9548E0312342018",
    "0x7Dd8a1037Faf2CE342c42ce43DBA9FB14Dd64a18",
    "0x8519d96C2d07753B26Ac399B477A396E34c443b2"
  ]
  },
  {
    "name": "XenoBunny",
    "slug": "xeno-bunny",
    "twitter": "@XenobunnyX",
    "logo": "xenobunny.jpg",
    "contracts": [
      "0x1996E10c64213Fe5E86AC7A7ac03Ec169176E4a7",
      "0xF645df6186AD0DbB088f7ef024C4Be640F1DaCd2",
      "0x553B1A22fB7c148690BC14014a9FDFc12e8Fdfa4"
    ]
  },
  {
    "name": "Swing",
    "slug": "swing",
    "logo": "swing.jpeg",
    "contracts": [
      "0x42df81c742CAe6F6D91E136b1AA5C7e14CB394FB",
      "0x90f1Ef9D2cDe204C8494Cf73130771B350070B53",
      "0x97fffFfa57144BBfacF41251bd1763657e646667"
    ]
  },
	{
		"name": "Symmetric",
		"slug": "symmetric",
    "contracts": [
			"0xbccc4b4c6530F82FE309c5E845E50b5E9C89f2AD",
			"0x4e4131dC27ed9501ac5fEb76F94572fDAe9f0fD0",
			"0xFEF39453770fF2C6b2F453D1b6D075623a79e3Eb",
			"0x030840b51ee6E96BA81423eF364234B8D656EAe8",
			"0x6cf35FD1c4649E94DFD11f5ab48e0a48D1211B3a",
			"0xdfFD6094ACa1A90f3064BB265e87245dD6f12e68",
			"0x0447cb41A2a0D4C4A76cA3e2b1be65076DD48A06",
			"0xc4E44B978c814F9223784031474ba1498bd23335",
			"0x76930FbaAbDB2D04B41835029D2320B2A0139cc5"
		]
	},
  {
    "name": "OpenALCHI",
    "slug": "openalchi",
    "contracts": [
      "0x9a92338b1CFfdf34b4Cc0192A8882F523b853e44",
      "0xed2147ff41e7B14a9C0c37fc4b7a1Eb7644EFb3C",
      "0x5cF295a3c53E090Cf230E1D40f413b8e3c225C26"
    ]
  },
  {
    "name": "Swapsicle",
    "slug": "swapsicle",
    "contracts": [
      "0xb68b27a1c93A52d698EecA5a759E2E4469432C09",
      "0xBa90FC740a95A6997306255853959Bb284cb748a",
      "0x7d027E68cD361F3A5F48ad52A09D21E10fD9AFAD",
      "0x2939E1c0DC976b4adE181CAFDfDC76F39f837F16",
      "0x6e497DbaC6a91CAeF4De3C6432838d25fC377a9e",
      "0x28B5556518b12b626Fe581fabF5C81Fb69f8e04A",
      "0xf93f75c7612f098098Fb2A511149F449225092f4",
      "0x3E016e10d04dD9Ce5E76F58f73B2ea0a1548e6Bc",
      "0x37F7B1187B758EA0655506A0a2cCdff523EaD9d5",
      "0x58596657e3d7FF8028Cb9597876e6e8674e61C96",
      "0x0cdde1dead51b156bd62113664d60b354b4df4ab",
      "0xe130d0dcafbd98f7f4ba357da7cd504212d6f0ce"
    ]
  },
  {
    "name": "Robinos",
    "slug": "robinos",
    "contracts": [
      "0x9EF9c57754eD079D750016b802DcCD45d0AB66f8"
    ]
  },
  {
    "name": "TakoTako",
    "slug": "takotako",
    "twitter": "@TAKOTAKOxyz",
    "logo": "takotako.png",
    "contracts": [
      "0x79a741EBFE9c323CF63180c405c050cdD98c21d8",
      "0x72C6bDf69952b6bc8aCc18c178d9E03EAc5eaD50",
      "0x7945F98240b310bD21F8814bdCEeBA6775a9A36A",
      "0x820C66D8316856655AdB42B3b6cB6a1728D29567",
      "0x6Afa285ab05657f7102F66F1B384347aEF3Ef6Aa",
      "0x19871b9911ddbd422e06F66427768f9B65d36F81",
      "0xbbFa45a92d9d071554B59D2d29174584D9b06bc3",
      "0x0f0244337f1215E6D8e13Af1b5ae639244d8a6f6"
    ]
  },
  {
    "name": "OmniHub",
    "slug": "omnihub",
    "twitter": "@omni_hub",
    "logo": "omnihub.jpg",
    "contracts": [
     "0xb0B4B761C9e9Bf5A9194a42e944A4A6646B83919"
    ]
  },
  {
    "name": "BarKrypto",
    "slug": "BarKrypto",
    "twitter": "@BarKrypto",
    "logo": "barkrypto.png",
    "contracts": [
      "0x72e5599798f36045c9549a11f2674d28c2209c30"
    ]
  }
]`
)
