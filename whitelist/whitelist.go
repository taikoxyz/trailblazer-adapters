package whitelist

type Protocol struct {
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	Contracts []string `json:"contracts"`
}

var (
	Contracts = `[
  {
    "name": "NFTs2Me",
    "slug": "nfts-2-me",
    "contracts": [
      "0x00000000001594C61dD8a6804da9AB58eD2483ce"
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
    "name": "Bunnyfi",
    "slug": "bunnyfi",
    "contracts": [
      "0x0000000000D310F8802cC91F198d14bC2303230B"
    ]
  },
  {
    "name": "Kodo Exchange",
    "slug": "kodo-exchange",
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
      "0x2f5c25874302558299e955d29619c6F4b9A16610",
      "0x64C310638a405D15A5C6C4590ac4A1f20d8A7546"
    ]
  },
  {
    "name": "Comet Protocol",
    "slug": "comet-protocol",
    "contracts": [
      "0xB50Ac92D6d8748AC42721c25A3e2C84637385A6b",
      "0x0fbCf4a62036E96C4F6770B38a9B536Aa14d1846"
    ]
  },
  {
    "name": "Stargate Finance",
    "slug": "stargate",
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
    "contracts": [
      "0x7ddB8A975778a434dE03dd666F11Ce962DCdD290",
      "0x2c301eBfB0bb42Af519377578099b63E921515B7",
      "0x6C8865042792B5E919fC95bF771ccaDF6F0cfA22",
      "0xD31A4be996b7E1cc20974181127E6fCA15649913",
      "0xA9EC1fEEE212851c829B028F094156CD04A3a547",
      "0x1ACa21A2a2a070d3536a69733c7044feDEB88f5A"
    ]
  },
  {
    "name": "HenjinDEX",
    "slug": "henjin",
    "contracts": [
      "0x07Bc9a408B385C7Aa8De2783795759512fE24356",
      "0xcFf128C67bCDc5a7c7D3F24c638e59AA0d4e112b",
      "0xaF0C5CBbCEfB685BF3200684d2aE19B8eA6186ca"
    ]
  },
  {
    "name": "Oku Trade",
    "slug": "oku",
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
    "contracts": [
      "0xFb2Cd41a8aeC89EFBb19575C6c48d872cE97A0A5"
    ]
  },
  {
    "name": "RubyScore",
    "slug": "ruby-score",
    "contracts": [
      "0xDC3D8318Fbaec2de49281843f5bba22e78338146",
      "0x4D1E2145082d0AB0fDa4a973dC4887C7295e21aB"
    ]
  },
  {
    "name": "Orbiter Finance",
    "slug": "orbiter",
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
    "contracts": [
      "0x1Df2De291F909baA50C1456C87C71Edf9Fb199D5"
    ]
  },
  {
    "name": "DTX.TRADE",
    "slug": "dtx",
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
    "contracts": [
      "0x009905bf008CcA637185EEaFE8F51BB56dD2ACa7"
    ]
  },
  {
    "name": "0xAstra",
    "slug": "0xastra",
    "contracts": [
      "0x90CE48ED68C6FCAe6F13b445F1573f003cF1804d",
      "0x34723B92aE9708BA33843120A86035D049dA7dfA"
    ]
  },
  {
    "name": "Ritsu Protocol",
    "slug": "ritsu",
    "contracts": [
      "0x7160570BB153Edd0Ea1775EC2b2Ac9b65F1aB61B",
      "0x12AF3Ec993EC5d5bD789b3e989c9E95A2F6c586D",
      "0xDFFee0ad5C833f2A5E610dfe9FD1aD82743eA74e",
      "0x3e846B1520E74728EFf445F1f86D348755F738d9"
    ]
  },
  {
    "name": "TaikoTown: A LooperLands Experience",
    "slug": "taiko-town",
    "contracts": [
      "0xEe01C4b0538849bF1c66bDFB458a7de11B1d7424"
    ]
  },
  {
    "name": "XY Finance",
    "slug": "xy-finance",
    "contracts": [
      "0x73Ce60416035B8D7019f6399778c14ccf5C9c7A1",
      "0xedC061306A79257f15108200C5B82ACc874C239d"
    ]
  },
  {
    "name": "KiloEx",
    "slug": "kiloex",
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
    "contracts": [
      "0xCA99F9DbF4A13D4de05B41a68041dcE7929cb5e0"
    ]
  },
  {
    "name": "Owlto Finance",
    "slug": "owlto-finance",
    "contracts": [
      "0x5e809A85Aa182A9921EDD10a4163745bb3e36284"
    ]
  },
  {
    "name": "coNFT",
    "slug": "co-nft",
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
    "contracts": [
      "0x00000000000007736e2F9aA5630B8c812E1F3fc9"
    ]
  },
  {
    "name": "Izumi Finance",
    "slug": "izumi-iziswap",
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
  }
]`
)
