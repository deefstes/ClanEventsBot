package main

import "strings"

// Emoji constants
const (
	EmojiGrinning                         = "😀"
	EmojiSmiley                           = "😃"
	EmojiSmile                            = "😄"
	EmojiGrin                             = "😁"
	EmojiLaughing                         = "😆"
	EmojiSweatSmile                       = "😅"
	EmojiJoy                              = "😂"
	EmojiRofl                             = "🤣"
	EmojiRelaxed                          = "☺️"
	EmojiBlush                            = "😊"
	EmojiInnocent                         = "😇"
	EmojiSlightlySmilingFace              = "🙂"
	EmojiUpsideDownFace                   = "🙃"
	EmojiWink                             = "😉"
	EmojiRelieved                         = "😌"
	EmojiHeartEyes                        = "😍"
	EmojiKissingHeart                     = "😘"
	EmojiKissing                          = "😗"
	EmojiKissingSmilingEyes               = "😙"
	EmojiKissingClosedEyes                = "😚"
	EmojiYum                              = "😋"
	EmojiStuckOutTongueWinkingEye         = "😜"
	EmojiStuckOutTongueClosedEyes         = "😝"
	EmojiStuckOutTongue                   = "😛"
	EmojiMoneyMouthFace                   = "🤑"
	EmojiHugs                             = "🤗"
	EmojiNerdFace                         = "🤓"
	EmojiSunglasses                       = "😎"
	EmojiClownFace                        = "🤡"
	EmojiCowboyHatFace                    = "🤠"
	EmojiSmirk                            = "😏"
	EmojiUnamused                         = "😒"
	EmojiDisappointed                     = "😞"
	EmojiPensive                          = "😔"
	EmojiWorried                          = "😟"
	EmojiConfused                         = "😕"
	EmojiSlightlyFrowningFace             = "🙁"
	EmojiFrowningFace                     = "☹️"
	EmojiPersevere                        = "😣"
	EmojiConfounded                       = "😖"
	EmojiTiredFace                        = "😫"
	EmojiWeary                            = "😩"
	EmojiTriumph                          = "😤"
	EmojiAngry                            = "😠"
	EmojiRage                             = "😡"
	EmojiNoMouth                          = "😶"
	EmojiNeutralFace                      = "😐"
	EmojiExpressionless                   = "😑"
	EmojiHushed                           = "😯"
	EmojiFrowning                         = "😦"
	EmojiAnguished                        = "😧"
	EmojiOpenMouth                        = "😮"
	EmojiAstonished                       = "😲"
	EmojiDizzyFace                        = "😵"
	EmojiFlushed                          = "😳"
	EmojiScream                           = "😱"
	EmojiFearful                          = "😨"
	EmojiColdSweat                        = "😰"
	EmojiCry                              = "😢"
	EmojiDisappointedRelieved             = "😥"
	EmojiDroolingFace                     = "🤤"
	EmojiSob                              = "😭"
	EmojiSweat                            = "😓"
	EmojiSleepy                           = "😪"
	EmojiSleeping                         = "😴"
	EmojiRollEyes                         = "🙄"
	EmojiThinking                         = "🤔"
	EmojiLyingFace                        = "🤥"
	EmojiGrimacing                        = "😬"
	EmojiZipperMouthFace                  = "🤐"
	EmojiNauseatedFace                    = "🤢"
	EmojiSneezingFace                     = "🤧"
	EmojiMask                             = "😷"
	EmojiFaceWithThermometer              = "🤒"
	EmojiFaceWithHeadBandage              = "🤕"
	EmojiSmilingImp                       = "😈"
	EmojiImp                              = "👿"
	EmojiJapaneseOgre                     = "👹"
	EmojiJapaneseGoblin                   = "👺"
	EmojiHankey                           = "💩"
	EmojiGhost                            = "👻"
	EmojiSkull                            = "💀"
	EmojiSkullAndCrossbones               = "☠️"
	EmojiAlien                            = "👽"
	EmojiSpaceInvader                     = "👾"
	EmojiRobot                            = "🤖"
	EmojiJackOLantern                     = "🎃"
	EmojiSmileyCat                        = "😺"
	EmojiSmileCat                         = "😸"
	EmojiJoyCat                           = "😹"
	EmojiHeartEyesCat                     = "😻"
	EmojiSmirkCat                         = "😼"
	EmojiKissingCat                       = "😽"
	EmojiScreamCat                        = "🙀"
	EmojiCryingCatFace                    = "😿"
	EmojiPoutingCat                       = "😾"
	EmojiOpenHands                        = "👐"
	EmojiRaisedHands                      = "🙌"
	EmojiClap                             = "👏"
	EmojiPray                             = "🙏"
	EmojiHandshake                        = "🤝"
	EmojiThumbsup                         = "👍"
	EmojiThumbsdown                       = "👎"
	EmojiFistOncoming                     = "👊"
	EmojiFistRaised                       = "✊"
	EmojiFistLeft                         = "🤛"
	EmojiFistRight                        = "🤜"
	EmojiCrossedFingers                   = "🤞"
	EmojiV                                = "✌️"
	EmojiMetal                            = "🤘"
	EmojiOkHand                           = "👌"
	EmojiPointLeft                        = "👈"
	EmojiPointRight                       = "👉"
	EmojiPointUp2                         = "👆"
	EmojiPointDown                        = "👇"
	EmojiPointUp                          = "☝️"
	EmojiHand                             = "✋"
	EmojiRaisedBackOfHand                 = "🤚"
	EmojiRaisedHandWithFingersSplayed     = "🖐"
	EmojiVulcanSalute                     = "🖖"
	EmojiWave                             = "👋"
	EmojiCallMeHand                       = "🤙"
	EmojiMuscle                           = "💪"
	EmojiMiddleFinger                     = "🖕"
	EmojiWritingHand                      = "✍️"
	EmojiSelfie                           = "🤳"
	EmojiNailCare                         = "💅"
	EmojiRing                             = "💍"
	EmojiLipstick                         = "💄"
	EmojiKiss                             = "💋"
	EmojiLips                             = "👄"
	EmojiTongue                           = "👅"
	EmojiEar                              = "👂"
	EmojiNose                             = "👃"
	EmojiFootprints                       = "👣"
	EmojiEye                              = "👁"
	EmojiEyes                             = "👀"
	EmojiSpeakingHead                     = "🗣"
	EmojiBustInSilhouette                 = "👤"
	EmojiBustsInSilhouette                = "👥"
	EmojiBaby                             = "👶"
	EmojiBoy                              = "👦"
	EmojiGirl                             = "👧"
	EmojiMan                              = "👨"
	EmojiWoman                            = "👩"
	EmojiBlondeWoman                      = "👱‍♀"
	EmojiBlondeMan                        = "👱"
	EmojiOlderMan                         = "👴"
	EmojiOlderWoman                       = "👵"
	EmojiManWithGuaPiMao                  = "👲"
	EmojiWomanWithTurban                  = "👳‍♀"
	EmojiManWithTurban                    = "👳"
	EmojiPolicewoman                      = "👮‍♀"
	EmojiPoliceman                        = "👮"
	EmojiConstructionWorkerWoman          = "👷‍♀"
	EmojiConstructionWorkerMan            = "👷"
	EmojiGuardswoman                      = "💂‍♀"
	EmojiGuardsman                        = "💂"
	EmojiFemaleDetective                  = "🕵️‍♀️"
	EmojiMaleDetective                    = "🕵"
	EmojiWomanHealthWorker                = "👩‍⚕"
	EmojiManHealthWorker                  = "👨‍⚕"
	EmojiWomanFarmer                      = "👩‍🌾"
	EmojiManFarmer                        = "👨‍🌾"
	EmojiWomanCook                        = "👩‍🍳"
	EmojiManCook                          = "👨‍🍳"
	EmojiWomanStudent                     = "👩‍🎓"
	EmojiManStudent                       = "👨‍🎓"
	EmojiWomanSinger                      = "👩‍🎤"
	EmojiManSinger                        = "👨‍🎤"
	EmojiWomanTeacher                     = "👩‍🏫"
	EmojiManTeacher                       = "👨‍🏫"
	EmojiWomanFactoryWorker               = "👩‍🏭"
	EmojiManFactoryWorker                 = "👨‍🏭"
	EmojiWomanTechnologist                = "👩‍💻"
	EmojiManTechnologist                  = "👨‍💻"
	EmojiWomanOfficeWorker                = "👩‍💼"
	EmojiManOfficeWorker                  = "👨‍💼"
	EmojiWomanMechanic                    = "👩‍🔧"
	EmojiManMechanic                      = "👨‍🔧"
	EmojiWomanScientist                   = "👩‍🔬"
	EmojiManScientist                     = "👨‍🔬"
	EmojiWomanArtist                      = "👩‍🎨"
	EmojiManArtist                        = "👨‍🎨"
	EmojiWomanFirefighter                 = "👩‍🚒"
	EmojiManFirefighter                   = "👨‍🚒"
	EmojiWomanPilot                       = "👩‍✈"
	EmojiManPilot                         = "👨‍✈"
	EmojiWomanAstronaut                   = "👩‍🚀"
	EmojiManAstronaut                     = "👨‍🚀"
	EmojiWomanJudge                       = "👩‍⚖"
	EmojiManJudge                         = "👨‍⚖"
	EmojiMrsClaus                         = "🤶"
	EmojiSanta                            = "🎅"
	EmojiPrincess                         = "👸"
	EmojiPrince                           = "🤴"
	EmojiBrideWithVeil                    = "👰"
	EmojiManInTuxedo                      = "🤵"
	EmojiAngel                            = "👼"
	EmojiPregnantWoman                    = "🤰"
	EmojiBowingWoman                      = "🙇‍♀"
	EmojiBowingMan                        = "🙇"
	EmojiTippingHandWoman                 = "💁"
	EmojiTippingHandMan                   = "💁‍♂"
	EmojiNoGoodWoman                      = "🙅"
	EmojiNoGoodMan                        = "🙅‍♂"
	EmojiOkWoman                          = "🙆"
	EmojiOkMan                            = "🙆‍♂"
	EmojiRaisingHandWoman                 = "🙋"
	EmojiRaisingHandMan                   = "🙋‍♂"
	EmojiWomanFacepalming                 = "🤦‍♀"
	EmojiManFacepalming                   = "🤦‍♂"
	EmojiWomanShrugging                   = "🤷‍♀"
	EmojiManShrugging                     = "🤷‍♂"
	EmojiPoutingWoman                     = "🙎"
	EmojiPoutingMan                       = "🙎‍♂"
	EmojiFrowningWoman                    = "🙍"
	EmojiFrowningMan                      = "🙍‍♂"
	EmojiHaircutWoman                     = "💇"
	EmojiHaircutMan                       = "💇‍♂"
	EmojiMassageWoman                     = "💆"
	EmojiMassageMan                       = "💆‍♂"
	EmojiBusinessSuitLevitating           = "🕴"
	EmojiDancer                           = "💃"
	EmojiManDancing                       = "🕺"
	EmojiDancingWomen                     = "👯"
	EmojiDancingMen                       = "👯‍♂"
	EmojiWalkingWoman                     = "🚶‍♀"
	EmojiWalkingMan                       = "🚶"
	EmojiRunningWoman                     = "🏃‍♀"
	EmojiRunningMan                       = "🏃"
	EmojiCouple                           = "👫"
	EmojiTwoWomenHoldingHands             = "👭"
	EmojiTwoMenHoldingHands               = "👬"
	EmojiCoupleWithHeartWomanMan          = "💑"
	EmojiCoupleWithHeartWomanWoman        = "👩‍❤️‍👩"
	EmojiCoupleWithHeartManMan            = "👨‍❤️‍👨"
	EmojiCouplekissManWoman               = "💏"
	EmojiCouplekissWomanWoman             = "👩‍❤️‍💋‍👩"
	EmojiCouplekissManMan                 = "👨‍❤️‍💋‍👨"
	EmojiFamilyManWomanBoy                = "👪"
	EmojiFamilyManWomanGirl               = "👨‍👩‍👧"
	EmojiFamilyManWomanGirlBoy            = "👨‍👩‍👧‍👦"
	EmojiFamilyManWomanBoyBoy             = "👨‍👩‍👦‍👦"
	EmojiFamilyManWomanGirlGirl           = "👨‍👩‍👧‍👧"
	EmojiFamilyWomanWomanBoy              = "👩‍👩‍👦"
	EmojiFamilyWomanWomanGirl             = "👩‍👩‍👧"
	EmojiFamilyWomanWomanGirlBoy          = "👩‍👩‍👧‍👦"
	EmojiFamilyWomanWomanBoyBoy           = "👩‍👩‍👦‍👦"
	EmojiFamilyWomanWomanGirlGirl         = "👩‍👩‍👧‍👧"
	EmojiFamilyManManBoy                  = "👨‍👨‍👦"
	EmojiFamilyManManGirl                 = "👨‍👨‍👧"
	EmojiFamilyManManGirlBoy              = "👨‍👨‍👧‍👦"
	EmojiFamilyManManBoyBoy               = "👨‍👨‍👦‍👦"
	EmojiFamilyManManGirlGirl             = "👨‍👨‍👧‍👧"
	EmojiFamilyWomanBoy                   = "👩‍👦"
	EmojiFamilyWomanGirl                  = "👩‍👧"
	EmojiFamilyWomanGirlBoy               = "👩‍👧‍👦"
	EmojiFamilyWomanBoyBoy                = "👩‍👦‍👦"
	EmojiFamilyWomanGirlGirl              = "👩‍👧‍👧"
	EmojiFamilyManBoy                     = "👨‍👦"
	EmojiFamilyManGirl                    = "👨‍👧"
	EmojiFamilyManGirlBoy                 = "👨‍👧‍👦"
	EmojiFamilyManBoyBoy                  = "👨‍👦‍👦"
	EmojiFamilyManGirlGirl                = "👨‍👧‍👧"
	EmojiWomansClothes                    = "👚"
	EmojiShirt                            = "👕"
	EmojiJeans                            = "👖"
	EmojiNecktie                          = "👔"
	EmojiDress                            = "👗"
	EmojiBikini                           = "👙"
	EmojiKimono                           = "👘"
	EmojiHighHeel                         = "👠"
	EmojiSandal                           = "👡"
	EmojiBoot                             = "👢"
	EmojiMansShoe                         = "👞"
	EmojiAthleticShoe                     = "👟"
	EmojiWomansHat                        = "👒"
	EmojiTophat                           = "🎩"
	EmojiMortarBoard                      = "🎓"
	EmojiCrown                            = "👑"
	EmojiRescueWorkerHelmet               = "⛑"
	EmojiSchoolSatchel                    = "🎒"
	EmojiPouch                            = "👝"
	EmojiPurse                            = "👛"
	EmojiHandbag                          = "👜"
	EmojiBriefcase                        = "💼"
	EmojiEyeglasses                       = "👓"
	EmojiDarkSunglasses                   = "🕶"
	EmojiClosedUmbrella                   = "🌂"
	EmojiOpenUmbrella                     = "☂️"
	EmojiDog                              = "🐶"
	EmojiCat                              = "🐱"
	EmojiMouse                            = "🐭"
	EmojiHamster                          = "🐹"
	EmojiRabbit                           = "🐰"
	EmojiFoxFace                          = "🦊"
	EmojiBear                             = "🐻"
	EmojiPandaFace                        = "🐼"
	EmojiKoala                            = "🐨"
	EmojiTiger                            = "🐯"
	EmojiLion                             = "🦁"
	EmojiCow                              = "🐮"
	EmojiPig                              = "🐷"
	EmojiPigNose                          = "🐽"
	EmojiFrog                             = "🐸"
	EmojiMonkeyFace                       = "🐵"
	EmojiSeeNoEvil                        = "🙈"
	EmojiHearNoEvil                       = "🙉"
	EmojiSpeakNoEvil                      = "🙊"
	EmojiMonkey                           = "🐒"
	EmojiChicken                          = "🐔"
	EmojiPenguin                          = "🐧"
	EmojiBird                             = "🐦"
	EmojiBabyChick                        = "🐤"
	EmojiHatchingChick                    = "🐣"
	EmojiHatchedChick                     = "🐥"
	EmojiDuck                             = "🦆"
	EmojiEagle                            = "🦅"
	EmojiOwl                              = "🦉"
	EmojiBat                              = "🦇"
	EmojiWolf                             = "🐺"
	EmojiBoar                             = "🐗"
	EmojiHorse                            = "🐴"
	EmojiUnicorn                          = "🦄"
	EmojiBee                              = "🐝"
	EmojiBug                              = "🐛"
	EmojiButterfly                        = "🦋"
	EmojiSnail                            = "🐌"
	EmojiShell                            = "🐚"
	EmojiBeetle                           = "🐞"
	EmojiAnt                              = "🐜"
	EmojiSpider                           = "🕷"
	EmojiSpiderWeb                        = "🕸"
	EmojiTurtle                           = "🐢"
	EmojiSnake                            = "🐍"
	EmojiLizard                           = "🦎"
	EmojiScorpion                         = "🦂"
	EmojiCrab                             = "🦀"
	EmojiSquid                            = "🦑"
	EmojiOctopus                          = "🐙"
	EmojiShrimp                           = "🦐"
	EmojiTropicalFish                     = "🐠"
	EmojiFish                             = "🐟"
	EmojiBlowfish                         = "🐡"
	EmojiDolphin                          = "🐬"
	EmojiShark                            = "🦈"
	EmojiWhale                            = "🐳"
	EmojiWhale2                           = "🐋"
	EmojiCrocodile                        = "🐊"
	EmojiLeopard                          = "🐆"
	EmojiTiger2                           = "🐅"
	EmojiWaterBuffalo                     = "🐃"
	EmojiOx                               = "🐂"
	EmojiCow2                             = "🐄"
	EmojiDeer                             = "🦌"
	EmojiDromedaryCamel                   = "🐪"
	EmojiCamel                            = "🐫"
	EmojiElephant                         = "🐘"
	EmojiRhinoceros                       = "🦏"
	EmojiGorilla                          = "🦍"
	EmojiRacehorse                        = "🐎"
	EmojiPig2                             = "🐖"
	EmojiGoat                             = "🐐"
	EmojiRam                              = "🐏"
	EmojiSheep                            = "🐑"
	EmojiDog2                             = "🐕"
	EmojiPoodle                           = "🐩"
	EmojiCat2                             = "🐈"
	EmojiRooster                          = "🐓"
	EmojiTurkey                           = "🦃"
	EmojiDove                             = "🕊"
	EmojiRabbit2                          = "🐇"
	EmojiMouse2                           = "🐁"
	EmojiRat                              = "🐀"
	EmojiChipmunk                         = "🐿"
	EmojiFeet                             = "🐾"
	EmojiDragon                           = "🐉"
	EmojiDragonFace                       = "🐲"
	EmojiCactus                           = "🌵"
	EmojiChristmasTree                    = "🎄"
	EmojiEvergreenTree                    = "🌲"
	EmojiDeciduousTree                    = "🌳"
	EmojiPalmTree                         = "🌴"
	EmojiSeedling                         = "🌱"
	EmojiHerb                             = "🌿"
	EmojiShamrock                         = "☘️"
	EmojiFourLeafClover                   = "🍀"
	EmojiBamboo                           = "🎍"
	EmojiTanabataTree                     = "🎋"
	EmojiLeaves                           = "🍃"
	EmojiFallenLeaf                       = "🍂"
	EmojiMapleLeaf                        = "🍁"
	EmojiMushroom                         = "🍄"
	EmojiEarOfRice                        = "🌾"
	EmojiBouquet                          = "💐"
	EmojiTulip                            = "🌷"
	EmojiRose                             = "🌹"
	EmojiWiltedFlower                     = "🥀"
	EmojiSunflower                        = "🌻"
	EmojiBlossom                          = "🌼"
	EmojiCherryBlossom                    = "🌸"
	EmojiHibiscus                         = "🌺"
	EmojiEarthAmericas                    = "🌎"
	EmojiEarthAfrica                      = "🌍"
	EmojiEarthAsia                        = "🌏"
	EmojiFullMoon                         = "🌕"
	EmojiWaningGibbousMoon                = "🌖"
	EmojiLastQuarterMoon                  = "🌗"
	EmojiWaningCrescentMoon               = "🌘"
	EmojiNewMoon                          = "🌑"
	EmojiWaxingCrescentMoon               = "🌒"
	EmojiFirstQuarterMoon                 = "🌓"
	EmojiMoon                             = "🌔"
	EmojiNewMoonWithFace                  = "🌚"
	EmojiFullMoonWithFace                 = "🌝"
	EmojiSunWithFace                      = "🌞"
	EmojiFirstQuarterMoonWithFace         = "🌛"
	EmojiLastQuarterMoonWithFace          = "🌜"
	EmojiCrescentMoon                     = "🌙"
	EmojiDizzy                            = "💫"
	EmojiStar                             = "⭐️"
	EmojiStar2                            = "🌟"
	EmojiSparkles                         = "✨"
	EmojiZap                              = "⚡️"
	EmojiFire                             = "🔥"
	EmojiBoom                             = "💥"
	EmojiComet                            = "☄"
	EmojiSunny                            = "☀️"
	EmojiSunBehindSmallCloud              = "🌤"
	EmojiPartlySunny                      = "⛅️"
	EmojiSunBehindLargeCloud              = "🌥"
	EmojiSunBehindRainCloud               = "🌦"
	EmojiRainbow                          = "🌈"
	EmojiCloud                            = "☁️"
	EmojiCloudWithRain                    = "🌧"
	EmojiCloudWithLightningAndRain        = "⛈"
	EmojiCloudWithLightning               = "🌩"
	EmojiCloudWithSnow                    = "🌨"
	EmojiSnowmanWithSnow                  = "☃️"
	EmojiSnowman                          = "⛄️"
	EmojiSnowflake                        = "❄️"
	EmojiWindFace                         = "🌬"
	EmojiDash                             = "💨"
	EmojiTornado                          = "🌪"
	EmojiFog                              = "🌫"
	EmojiOcean                            = "🌊"
	EmojiDroplet                          = "💧"
	EmojiSweatDrops                       = "💦"
	EmojiUmbrella                         = "☔️"
	EmojiGreenApple                       = "🍏"
	EmojiApple                            = "🍎"
	EmojiPear                             = "🍐"
	EmojiTangerine                        = "🍊"
	EmojiLemon                            = "🍋"
	EmojiBanana                           = "🍌"
	EmojiWatermelon                       = "🍉"
	EmojiGrapes                           = "🍇"
	EmojiStrawberry                       = "🍓"
	EmojiMelon                            = "🍈"
	EmojiCherries                         = "🍒"
	EmojiPeach                            = "🍑"
	EmojiPineapple                        = "🍍"
	EmojiKiwiFruit                        = "🥝"
	EmojiAvocado                          = "🥑"
	EmojiTomato                           = "🍅"
	EmojiEggplant                         = "🍆"
	EmojiCucumber                         = "🥒"
	EmojiCarrot                           = "🥕"
	EmojiCorn                             = "🌽"
	EmojiHotPepper                        = "🌶"
	EmojiPotato                           = "🥔"
	EmojiSweetPotato                      = "🍠"
	EmojiChestnut                         = "🌰"
	EmojiPeanuts                          = "🥜"
	EmojiHoneyPot                         = "🍯"
	EmojiCroissant                        = "🥐"
	EmojiBread                            = "🍞"
	EmojiBaguetteBread                    = "🥖"
	EmojiCheese                           = "🧀"
	EmojiEgg                              = "🥚"
	EmojiFriedEgg                         = "🍳"
	EmojiBacon                            = "🥓"
	EmojiPancakes                         = "🥞"
	EmojiFriedShrimp                      = "🍤"
	EmojiPoultryLeg                       = "🍗"
	EmojiMeatOnBone                       = "🍖"
	EmojiPizza                            = "🍕"
	EmojiHotdog                           = "🌭"
	EmojiHamburger                        = "🍔"
	EmojiFries                            = "🍟"
	EmojiStuffedFlatbread                 = "🥙"
	EmojiTaco                             = "🌮"
	EmojiBurrito                          = "🌯"
	EmojiGreenSalad                       = "🥗"
	EmojiShallowPanOfFood                 = "🥘"
	EmojiSpaghetti                        = "🍝"
	EmojiRamen                            = "🍜"
	EmojiStew                             = "🍲"
	EmojiFishCake                         = "🍥"
	EmojiSushi                            = "🍣"
	EmojiBento                            = "🍱"
	EmojiCurry                            = "🍛"
	EmojiRice                             = "🍚"
	EmojiRiceBall                         = "🍙"
	EmojiRiceCracker                      = "🍘"
	EmojiOden                             = "🍢"
	EmojiDango                            = "🍡"
	EmojiShavedIce                        = "🍧"
	EmojiIceCream                         = "🍨"
	EmojiIcecream                         = "🍦"
	EmojiCake                             = "🍰"
	EmojiBirthday                         = "🎂"
	EmojiCustard                          = "🍮"
	EmojiLollipop                         = "🍭"
	EmojiCandy                            = "🍬"
	EmojiChocolateBar                     = "🍫"
	EmojiPopcorn                          = "🍿"
	EmojiDoughnut                         = "🍩"
	EmojiCookie                           = "🍪"
	EmojiMilkGlass                        = "🥛"
	EmojiBabyBottle                       = "🍼"
	EmojiCoffee                           = "☕️"
	EmojiTea                              = "🍵"
	EmojiSake                             = "🍶"
	EmojiBeer                             = "🍺"
	EmojiBeers                            = "🍻"
	EmojiClinkingGlasses                  = "🥂"
	EmojiWineGlass                        = "🍷"
	EmojiTumblerGlass                     = "🥃"
	EmojiCocktail                         = "🍸"
	EmojiTropicalDrink                    = "🍹"
	EmojiChampagne                        = "🍾"
	EmojiSpoon                            = "🥄"
	EmojiForkAndKnife                     = "🍴"
	EmojiPlateWithCutlery                 = "🍽"
	EmojiSoccer                           = "⚽️"
	EmojiBasketball                       = "🏀"
	EmojiFootball                         = "🏈"
	EmojiBaseball                         = "⚾️"
	EmojiTennis                           = "🎾"
	EmojiVolleyball                       = "🏐"
	EmojiRugbyFootball                    = "🏉"
	Emoji8ball                            = "🎱"
	EmojiPingPong                         = "🏓"
	EmojiBadminton                        = "🏸"
	EmojiGoalNet                          = "🥅"
	EmojiIceHockey                        = "🏒"
	EmojiFieldHockey                      = "🏑"
	EmojiCricket                          = "🏏"
	EmojiGolf                             = "⛳️"
	EmojiBowAndArrow                      = "🏹"
	EmojiFishingPoleAndFish               = "🎣"
	EmojiBoxingGlove                      = "🥊"
	EmojiMartialArtsUniform               = "🥋"
	EmojiIceSkate                         = "⛸"
	EmojiSki                              = "🎿"
	EmojiSkier                            = "⛷"
	EmojiSnowboarder                      = "🏂"
	EmojiWeightLiftingWoman               = "🏋️‍♀️"
	EmojiWeightLiftingMan                 = "🏋"
	EmojiPersonFencing                    = "🤺"
	EmojiWomenWrestling                   = "🤼‍♀"
	EmojiMenWrestling                     = "🤼‍♂"
	EmojiWomanCartwheeling                = "🤸‍♀"
	EmojiManCartwheeling                  = "🤸‍♂"
	EmojiBasketballWoman                  = "⛹️‍♀️"
	EmojiBasketballMan                    = "⛹"
	EmojiWomanPlayingHandball             = "🤾‍♀"
	EmojiManPlayingHandball               = "🤾‍♂"
	EmojiGolfingWoman                     = "🏌️‍♀️"
	EmojiGolfingMan                       = "🏌"
	EmojiSurfingWoman                     = "🏄‍♀"
	EmojiSurfingMan                       = "🏄"
	EmojiSwimmingWoman                    = "🏊‍♀"
	EmojiSwimmingMan                      = "🏊"
	EmojiWomanPlayingWaterPolo            = "🤽‍♀"
	EmojiManPlayingWaterPolo              = "🤽‍♂"
	EmojiRowingWoman                      = "🚣‍♀"
	EmojiRowingMan                        = "🚣"
	EmojiHorseRacing                      = "🏇"
	EmojiBikingWoman                      = "🚴‍♀"
	EmojiBikingMan                        = "🚴"
	EmojiMountainBikingWoman              = "🚵‍♀"
	EmojiMountainBikingMan                = "🚵"
	EmojiRunningShirtWithSash             = "🎽"
	EmojiMedalSports                      = "🏅"
	EmojiMedalMilitary                    = "🎖"
	Emoji1stPlaceMedal                    = "🥇"
	Emoji2ndPlaceMedal                    = "🥈"
	Emoji3rdPlaceMedal                    = "🥉"
	EmojiTrophy                           = "🏆"
	EmojiRosette                          = "🏵"
	EmojiReminderRibbon                   = "🎗"
	EmojiTicket                           = "🎫"
	EmojiTickets                          = "🎟"
	EmojiCircusTent                       = "🎪"
	EmojiWomanJuggling                    = "🤹‍♀"
	EmojiManJuggling                      = "🤹‍♂"
	EmojiPerformingArts                   = "🎭"
	EmojiArt                              = "🎨"
	EmojiClapper                          = "🎬"
	EmojiMicrophone                       = "🎤"
	EmojiHeadphones                       = "🎧"
	EmojiMusicalScore                     = "🎼"
	EmojiMusicalKeyboard                  = "🎹"
	EmojiDrum                             = "🥁"
	EmojiSaxophone                        = "🎷"
	EmojiTrumpet                          = "🎺"
	EmojiGuitar                           = "🎸"
	EmojiViolin                           = "🎻"
	EmojiGameDie                          = "🎲"
	EmojiDart                             = "🎯"
	EmojiBowling                          = "🎳"
	EmojiVideoGame                        = "🎮"
	EmojiSlotMachine                      = "🎰"
	EmojiCar                              = "🚗"
	EmojiTaxi                             = "🚕"
	EmojiBlueCar                          = "🚙"
	EmojiBus                              = "🚌"
	EmojiTrolleybus                       = "🚎"
	EmojiRacingCar                        = "🏎"
	EmojiPoliceCar                        = "🚓"
	EmojiAmbulance                        = "🚑"
	EmojiFireEngine                       = "🚒"
	EmojiMinibus                          = "🚐"
	EmojiTruck                            = "🚚"
	EmojiArticulatedLorry                 = "🚛"
	EmojiTractor                          = "🚜"
	EmojiKickScooter                      = "🛴"
	EmojiBike                             = "🚲"
	EmojiMotorScooter                     = "🛵"
	EmojiMotorcycle                       = "🏍"
	EmojiRotatingLight                    = "🚨"
	EmojiOncomingPoliceCar                = "🚔"
	EmojiOncomingBus                      = "🚍"
	EmojiOncomingAutomobile               = "🚘"
	EmojiOncomingTaxi                     = "🚖"
	EmojiAerialTramway                    = "🚡"
	EmojiMountainCableway                 = "🚠"
	EmojiSuspensionRailway                = "🚟"
	EmojiRailwayCar                       = "🚃"
	EmojiTrain                            = "🚋"
	EmojiMountainRailway                  = "🚞"
	EmojiMonorail                         = "🚝"
	EmojiBullettrainSide                  = "🚄"
	EmojiBullettrainFront                 = "🚅"
	EmojiLightRail                        = "🚈"
	EmojiSteamLocomotive                  = "🚂"
	EmojiTrain2                           = "🚆"
	EmojiMetro                            = "🚇"
	EmojiTram                             = "🚊"
	EmojiStation                          = "🚉"
	EmojiHelicopter                       = "🚁"
	EmojiSmallAirplane                    = "🛩"
	EmojiAirplane                         = "✈️"
	EmojiFlightDeparture                  = "🛫"
	EmojiFlightArrival                    = "🛬"
	EmojiRocket                           = "🚀"
	EmojiArtificialSatellite              = "🛰"
	EmojiSeat                             = "💺"
	EmojiCanoe                            = "🛶"
	EmojiBoat                             = "⛵️"
	EmojiMotorBoat                        = "🛥"
	EmojiSpeedboat                        = "🚤"
	EmojiPassengerShip                    = "🛳"
	EmojiFerry                            = "⛴"
	EmojiShip                             = "🚢"
	EmojiAnchor                           = "⚓️"
	EmojiConstruction                     = "🚧"
	EmojiFuelpump                         = "⛽️"
	EmojiBusstop                          = "🚏"
	EmojiVerticalTrafficLight             = "🚦"
	EmojiTrafficLight                     = "🚥"
	EmojiWorldMap                         = "🗺"
	EmojiMoyai                            = "🗿"
	EmojiStatueOfLiberty                  = "🗽"
	EmojiFountain                         = "⛲️"
	EmojiTokyoTower                       = "🗼"
	EmojiEuropeanCastle                   = "🏰"
	EmojiJapaneseCastle                   = "🏯"
	EmojiStadium                          = "🏟"
	EmojiFerrisWheel                      = "🎡"
	EmojiRollerCoaster                    = "🎢"
	EmojiCarouselHorse                    = "🎠"
	EmojiParasolOnGround                  = "⛱"
	EmojiBeachUmbrella                    = "🏖"
	EmojiDesertIsland                     = "🏝"
	EmojiMountain                         = "⛰"
	EmojiMountainSnow                     = "🏔"
	EmojiMountFuji                        = "🗻"
	EmojiVolcano                          = "🌋"
	EmojiDesert                           = "🏜"
	EmojiCamping                          = "🏕"
	EmojiTent                             = "⛺️"
	EmojiRailwayTrack                     = "🛤"
	EmojiMotorway                         = "🛣"
	EmojiBuildingConstruction             = "🏗"
	EmojiFactory                          = "🏭"
	EmojiHouse                            = "🏠"
	EmojiHouseWithGarden                  = "🏡"
	EmojiHouses                           = "🏘"
	EmojiDerelictHouse                    = "🏚"
	EmojiOffice                           = "🏢"
	EmojiDepartmentStore                  = "🏬"
	EmojiPostOffice                       = "🏣"
	EmojiEuropeanPostOffice               = "🏤"
	EmojiHospital                         = "🏥"
	EmojiBank                             = "🏦"
	EmojiHotel                            = "🏨"
	EmojiConvenienceStore                 = "🏪"
	EmojiSchool                           = "🏫"
	EmojiLoveHotel                        = "🏩"
	EmojiWedding                          = "💒"
	EmojiClassicalBuilding                = "🏛"
	EmojiChurch                           = "⛪️"
	EmojiMosque                           = "🕌"
	EmojiSynagogue                        = "🕍"
	EmojiKaaba                            = "🕋"
	EmojiShintoShrine                     = "⛩"
	EmojiJapan                            = "🗾"
	EmojiRiceScene                        = "🎑"
	EmojiNationalPark                     = "🏞"
	EmojiSunrise                          = "🌅"
	EmojiSunriseOverMountains             = "🌄"
	EmojiStars                            = "🌠"
	EmojiSparkler                         = "🎇"
	EmojiFireworks                        = "🎆"
	EmojiCitySunrise                      = "🌇"
	EmojiCitySunset                       = "🌆"
	EmojiCityscape                        = "🏙"
	EmojiNightWithStars                   = "🌃"
	EmojiMilkyWay                         = "🌌"
	EmojiBridgeAtNight                    = "🌉"
	EmojiFoggy                            = "🌁"
	EmojiWatch                            = "⌚️"
	EmojiIphone                           = "📱"
	EmojiCalling                          = "📲"
	EmojiComputer                         = "💻"
	EmojiKeyboard                         = "⌨️"
	EmojiDesktopComputer                  = "🖥"
	EmojiPrinter                          = "🖨"
	EmojiComputerMouse                    = "🖱"
	EmojiTrackball                        = "🖲"
	EmojiJoystick                         = "🕹"
	EmojiClamp                            = "🗜"
	EmojiMinidisc                         = "💽"
	EmojiFloppyDisk                       = "💾"
	EmojiCd                               = "💿"
	EmojiDvd                              = "📀"
	EmojiVhs                              = "📼"
	EmojiCamera                           = "📷"
	EmojiCameraFlash                      = "📸"
	EmojiVideoCamera                      = "📹"
	EmojiMovieCamera                      = "🎥"
	EmojiFilmProjector                    = "📽"
	EmojiFilmStrip                        = "🎞"
	EmojiTelephoneReceiver                = "📞"
	EmojiPhone                            = "☎️"
	EmojiPager                            = "📟"
	EmojiFax                              = "📠"
	EmojiTv                               = "📺"
	EmojiRadio                            = "📻"
	EmojiStudioMicrophone                 = "🎙"
	EmojiLevelSlider                      = "🎚"
	EmojiControlKnobs                     = "🎛"
	EmojiStopwatch                        = "⏱"
	EmojiTimerClock                       = "⏲"
	EmojiAlarmClock                       = "⏰"
	EmojiMantelpieceClock                 = "🕰"
	EmojiHourglass                        = "⌛️"
	EmojiHourglassFlowingSand             = "⏳"
	EmojiSatellite                        = "📡"
	EmojiBattery                          = "🔋"
	EmojiElectricPlug                     = "🔌"
	EmojiBulb                             = "💡"
	EmojiFlashlight                       = "🔦"
	EmojiCandle                           = "🕯"
	EmojiWastebasket                      = "🗑"
	EmojiOilDrum                          = "🛢"
	EmojiMoneyWithWings                   = "💸"
	EmojiDollar                           = "💵"
	EmojiYen                              = "💴"
	EmojiEuro                             = "💶"
	EmojiPound                            = "💷"
	EmojiMoneybag                         = "💰"
	EmojiCreditCard                       = "💳"
	EmojiGem                              = "💎"
	EmojiBalanceScale                     = "⚖️"
	EmojiWrench                           = "🔧"
	EmojiHammer                           = "🔨"
	EmojiHammerAndPick                    = "⚒"
	EmojiHammerAndWrench                  = "🛠"
	EmojiPick                             = "⛏"
	EmojiNutAndBolt                       = "🔩"
	EmojiGear                             = "⚙️"
	EmojiChains                           = "⛓"
	EmojiGun                              = "🔫"
	EmojiBomb                             = "💣"
	EmojiHocho                            = "🔪"
	EmojiDagger                           = "🗡"
	EmojiCrossedSwords                    = "⚔️"
	EmojiShield                           = "🛡"
	EmojiSmoking                          = "🚬"
	EmojiCoffin                           = "⚰️"
	EmojiFuneralUrn                       = "⚱️"
	EmojiAmphora                          = "🏺"
	EmojiCrystalBall                      = "🔮"
	EmojiPrayerBeads                      = "📿"
	EmojiBarber                           = "💈"
	EmojiAlembic                          = "⚗️"
	EmojiTelescope                        = "🔭"
	EmojiMicroscope                       = "🔬"
	EmojiHole                             = "🕳"
	EmojiPill                             = "💊"
	EmojiSyringe                          = "💉"
	EmojiThermometer                      = "🌡"
	EmojiToilet                           = "🚽"
	EmojiPotableWater                     = "🚰"
	EmojiShower                           = "🚿"
	EmojiBathtub                          = "🛁"
	EmojiBath                             = "🛀"
	EmojiBellhopBell                      = "🛎"
	EmojiKey                              = "🔑"
	EmojiOldKey                           = "🗝"
	EmojiDoor                             = "🚪"
	EmojiCouchAndLamp                     = "🛋"
	EmojiBed                              = "🛏"
	EmojiSleepingBed                      = "🛌"
	EmojiFramedPicture                    = "🖼"
	EmojiShopping                         = "🛍"
	EmojiShoppingCart                     = "🛒"
	EmojiGift                             = "🎁"
	EmojiBalloon                          = "🎈"
	EmojiFlags                            = "🎏"
	EmojiRibbon                           = "🎀"
	EmojiConfettiBall                     = "🎊"
	EmojiTada                             = "🎉"
	EmojiDolls                            = "🎎"
	EmojiIzakayaLantern                   = "🏮"
	EmojiWindChime                        = "🎐"
	EmojiEmail                            = "✉️"
	EmojiEnvelopeWithArrow                = "📩"
	EmojiIncomingEnvelope                 = "📨"
	EmojiLoveLetter                       = "💌"
	EmojiInboxTray                        = "📥"
	EmojiOutboxTray                       = "📤"
	EmojiPackage                          = "📦"
	EmojiLabel                            = "🏷"
	EmojiMailboxClosed                    = "📪"
	EmojiMailbox                          = "📫"
	EmojiMailboxWithMail                  = "📬"
	EmojiMailboxWithNoMail                = "📭"
	EmojiPostbox                          = "📮"
	EmojiPostalHorn                       = "📯"
	EmojiScroll                           = "📜"
	EmojiPageWithCurl                     = "📃"
	EmojiPageFacingUp                     = "📄"
	EmojiBookmarkTabs                     = "📑"
	EmojiBarChart                         = "📊"
	EmojiChartWithUpwardsTrend            = "📈"
	EmojiChartWithDownwardsTrend          = "📉"
	EmojiSpiralNotepad                    = "🗒"
	EmojiSpiralCalendar                   = "🗓"
	EmojiCalendar                         = "📆"
	EmojiDate                             = "📅"
	EmojiCardIndex                        = "📇"
	EmojiCardFileBox                      = "🗃"
	EmojiBallotBox                        = "🗳"
	EmojiFileCabinet                      = "🗄"
	EmojiClipboard                        = "📋"
	EmojiFileFolder                       = "📁"
	EmojiOpenFileFolder                   = "📂"
	EmojiCardIndexDividers                = "🗂"
	EmojiNewspaperRoll                    = "🗞"
	EmojiNewspaper                        = "📰"
	EmojiNotebook                         = "📓"
	EmojiNotebookWithDecorativeCover      = "📔"
	EmojiLedger                           = "📒"
	EmojiClosedBook                       = "📕"
	EmojiGreenBook                        = "📗"
	EmojiBlueBook                         = "📘"
	EmojiOrangeBook                       = "📙"
	EmojiBooks                            = "📚"
	EmojiBook                             = "📖"
	EmojiBookmark                         = "🔖"
	EmojiLink                             = "🔗"
	EmojiPaperclip                        = "📎"
	EmojiPaperclips                       = "🖇"
	EmojiTriangularRuler                  = "📐"
	EmojiStraightRuler                    = "📏"
	EmojiPushpin                          = "📌"
	EmojiRoundPushpin                     = "📍"
	EmojiScissors                         = "✂️"
	EmojiPen                              = "🖊"
	EmojiFountainPen                      = "🖋"
	EmojiBlackNib                         = "✒️"
	EmojiPaintbrush                       = "🖌"
	EmojiCrayon                           = "🖍"
	EmojiMemo                             = "📝"
	EmojiPencil2                          = "✏️"
	EmojiMag                              = "🔍"
	EmojiMagRight                         = "🔎"
	EmojiLockWithInkPen                   = "🔏"
	EmojiClosedLockWithKey                = "🔐"
	EmojiLock                             = "🔒"
	EmojiUnlock                           = "🔓"
	EmojiHeart                            = "❤️"
	EmojiYellowHeart                      = "💛"
	EmojiGreenHeart                       = "💚"
	EmojiBlueHeart                        = "💙"
	EmojiPurpleHeart                      = "💜"
	EmojiBlackHeart                       = "🖤"
	EmojiBrokenHeart                      = "💔"
	EmojiHeavyHeartExclamation            = "❣️"
	EmojiTwoHearts                        = "💕"
	EmojiRevolvingHearts                  = "💞"
	EmojiHeartbeat                        = "💓"
	EmojiHeartpulse                       = "💗"
	EmojiSparklingHeart                   = "💖"
	EmojiCupid                            = "💘"
	EmojiGiftHeart                        = "💝"
	EmojiHeartDecoration                  = "💟"
	EmojiPeaceSymbol                      = "☮️"
	EmojiLatinCross                       = "✝️"
	EmojiStarAndCrescent                  = "☪️"
	EmojiOm                               = "🕉"
	EmojiWheelOfDharma                    = "☸️"
	EmojiStarOfDavid                      = "✡️"
	EmojiSixPointedStar                   = "🔯"
	EmojiMenorah                          = "🕎"
	EmojiYinYang                          = "☯️"
	EmojiOrthodoxCross                    = "☦️"
	EmojiPlaceOfWorship                   = "🛐"
	EmojiOphiuchus                        = "⛎"
	EmojiAries                            = "♈️"
	EmojiTaurus                           = "♉️"
	EmojiGemini                           = "♊️"
	EmojiCancer                           = "♋️"
	EmojiLeo                              = "♌️"
	EmojiVirgo                            = "♍️"
	EmojiLibra                            = "♎️"
	EmojiScorpius                         = "♏️"
	EmojiSagittarius                      = "♐️"
	EmojiCapricorn                        = "♑️"
	EmojiAquarius                         = "♒️"
	EmojiPisces                           = "♓️"
	EmojiId                               = "🆔"
	EmojiAtomSymbol                       = "⚛️"
	EmojiAccept                           = "🉑"
	EmojiRadioactive                      = "☢️"
	EmojiBiohazard                        = "☣️"
	EmojiMobilePhoneOff                   = "📴"
	EmojiVibrationMode                    = "📳"
	EmojiU6709                            = "🈶"
	EmojiU7121                            = "🈚️"
	EmojiU7533                            = "🈸"
	EmojiU55b6                            = "🈺"
	EmojiU6708                            = "🈷️"
	EmojiEightPointedBlackStar            = "✴️"
	EmojiVs                               = "🆚"
	EmojiWhiteFlower                      = "💮"
	EmojiIdeographAdvantage               = "🉐"
	EmojiSecret                           = "㊙️"
	EmojiCongratulations                  = "㊗️"
	EmojiU5408                            = "🈴"
	EmojiU6e80                            = "🈵"
	EmojiU5272                            = "🈹"
	EmojiU7981                            = "🈲"
	EmojiA                                = "🅰️"
	EmojiB                                = "🅱️"
	EmojiAb                               = "🆎"
	EmojiCl                               = "🆑"
	EmojiO2                               = "🅾️"
	EmojiSos                              = "🆘"
	EmojiX                                = "❌"
	EmojiO                                = "⭕️"
	EmojiStopSign                         = "🛑"
	EmojiNoEntry                          = "⛔️"
	EmojiNameBadge                        = "📛"
	EmojiNoEntrySign                      = "🚫"
	Emoji100                              = "💯"
	EmojiAnger                            = "💢"
	EmojiHotsprings                       = "♨️"
	EmojiNoPedestrians                    = "🚷"
	EmojiDoNotLitter                      = "🚯"
	EmojiNoBicycles                       = "🚳"
	EmojiUnderage                         = "🔞"
	EmojiNoMobilePhones                   = "📵"
	EmojiNoSmoking                        = "🚭"
	EmojiExclamation                      = "❗️"
	EmojiGreyExclamation                  = "❕"
	EmojiQuestion                         = "❓"
	EmojiGreyQuestion                     = "❔"
	EmojiBangbang                         = "‼️"
	EmojiInterrobang                      = "⁉️"
	EmojiLowBrightness                    = "🔅"
	EmojiHighBrightness                   = "🔆"
	EmojiPartAlternationMark              = "〽️"
	EmojiWarning                          = "⚠️"
	EmojiChildrenCrossing                 = "🚸"
	EmojiTrident                          = "🔱"
	EmojiFleurDeLis                       = "⚜️"
	EmojiBeginner                         = "🔰"
	EmojiRecycle                          = "♻️"
	EmojiWhiteCheckMark                   = "✅"
	EmojiU6307                            = "🈯️"
	EmojiChart                            = "💹"
	EmojiSparkle                          = "❇️"
	EmojiEightSpokedAsterisk              = "✳️"
	EmojiNegativeSquaredCrossMark         = "❎"
	EmojiGlobeWithMeridians               = "🌐"
	EmojiDiamondShapeWithADotInside       = "💠"
	EmojiM                                = "Ⓜ️"
	EmojiCyclone                          = "🌀"
	EmojiZzz                              = "💤"
	EmojiAtm                              = "🏧"
	EmojiWc                               = "🚾"
	EmojiWheelchair                       = "♿️"
	EmojiParking                          = "🅿️"
	EmojiU7a7a                            = "🈳"
	EmojiSa                               = "🈂️"
	EmojiPassportControl                  = "🛂"
	EmojiCustoms                          = "🛃"
	EmojiBaggageClaim                     = "🛄"
	EmojiLeftLuggage                      = "🛅"
	EmojiMens                             = "🚹"
	EmojiWomens                           = "🚺"
	EmojiBabySymbol                       = "🚼"
	EmojiRestroom                         = "🚻"
	EmojiPutLitterInItsPlace              = "🚮"
	EmojiCinema                           = "🎦"
	EmojiSignalStrength                   = "📶"
	EmojiKoko                             = "🈁"
	EmojiSymbols                          = "🔣"
	EmojiInformationSource                = "ℹ️"
	EmojiAbc                              = "🔤"
	EmojiAbcd                             = "🔡"
	EmojiCapitalAbcd                      = "🔠"
	EmojiNg                               = "🆖"
	EmojiOk                               = "🆗"
	EmojiUp                               = "🆙"
	EmojiCool                             = "🆒"
	EmojiNew                              = "🆕"
	EmojiFree                             = "🆓"
	EmojiZero                             = "\x30\xe2\x83\xa3"
	EmojiOne                              = "\x31\xe2\x83\xa3"
	EmojiTwo                              = "\x32\xe2\x83\xa3"
	EmojiThree                            = "\x33\xe2\x83\xa3"
	EmojiFour                             = "\x34\xe2\x83\xa3"
	EmojiFive                             = "\x35\xe2\x83\xa3"
	EmojiSix                              = "\x36\xe2\x83\xa3"
	EmojiSeven                            = "\x37\xe2\x83\xa3"
	EmojiEight                            = "\x38\xe2\x83\xa3"
	EmojiNine                             = "\x39\xe2\x83\xa3"
	EmojiKeycapTen                        = "🔟"
	Emoji1234                             = "🔢"
	EmojiHash                             = "#️⃣"
	EmojiAsterisk                         = "*️⃣"
	EmojiArrowForward                     = "▶️"
	EmojiPauseButton                      = "⏸"
	EmojiPlayOrPauseButton                = "⏯"
	EmojiStopButton                       = "⏹"
	EmojiRecordButton                     = "⏺"
	EmojiNextTrackButton                  = "⏭"
	EmojiPreviousTrackButton              = "⏮"
	EmojiFastForward                      = "⏩"
	EmojiRewind                           = "⏪"
	EmojiArrowDoubleUp                    = "⏫"
	EmojiArrowDoubleDown                  = "⏬"
	EmojiArrowBackward                    = "◀️"
	EmojiArrowUpSmall                     = "🔼"
	EmojiArrowDownSmall                   = "🔽"
	EmojiArrowRight                       = "➡️"
	EmojiArrowLeft                        = "⬅️"
	EmojiArrowUp                          = "⬆️"
	EmojiArrowDown                        = "⬇️"
	EmojiArrowUpperRight                  = "↗️"
	EmojiArrowLowerRight                  = "↘️"
	EmojiArrowLowerLeft                   = "↙️"
	EmojiArrowUpperLeft                   = "↖️"
	EmojiArrowUpDown                      = "↕️"
	EmojiLeftRightArrow                   = "↔️"
	EmojiArrowRightHook                   = "↪️"
	EmojiLeftwardsArrowWithHook           = "↩️"
	EmojiArrowHeadingUp                   = "⤴️"
	EmojiArrowHeadingDown                 = "⤵️"
	EmojiTwistedRightwardsArrows          = "🔀"
	EmojiRepeat                           = "🔁"
	EmojiRepeatOne                        = "🔂"
	EmojiArrowsCounterclockwise           = "🔄"
	EmojiArrowsClockwise                  = "🔃"
	EmojiMusicalNote                      = "🎵"
	EmojiNotes                            = "🎶"
	EmojiHeavyPlusSign                    = "➕"
	EmojiHeavyMinusSign                   = "➖"
	EmojiHeavyDivisionSign                = "➗"
	EmojiHeavyMultiplicationX             = "✖️"
	EmojiHeavyDollarSign                  = "💲"
	EmojiCurrencyExchange                 = "💱"
	EmojiTm                               = "™️"
	EmojiCopyright                        = "©️"
	EmojiRegistered                       = "®️"
	EmojiWavyDash                         = "〰️"
	EmojiCurlyLoop                        = "➰"
	EmojiLoop                             = "➿"
	EmojiEnd                              = "🔚"
	EmojiBack                             = "🔙"
	EmojiOn                               = "🔛"
	EmojiTop                              = "🔝"
	EmojiSoon                             = "🔜"
	EmojiHeavyCheckMark                   = "✔️"
	EmojiBallotBoxWithCheck               = "☑️"
	EmojiRadioButton                      = "🔘"
	EmojiWhiteCircle                      = "⚪️"
	EmojiBlackCircle                      = "⚫️"
	EmojiRedCircle                        = "🔴"
	EmojiLargeBlueCircle                  = "🔵"
	EmojiSmallRedTriangle                 = "🔺"
	EmojiSmallRedTriangleDown             = "🔻"
	EmojiSmallOrangeDiamond               = "🔸"
	EmojiSmallBlueDiamond                 = "🔹"
	EmojiLargeOrangeDiamond               = "🔶"
	EmojiLargeBlueDiamond                 = "🔷"
	EmojiWhiteSquareButton                = "🔳"
	EmojiBlackSquareButton                = "🔲"
	EmojiBlackSmallSquare                 = "▪️"
	EmojiWhiteSmallSquare                 = "▫️"
	EmojiBlackMediumSmallSquare           = "◾️"
	EmojiWhiteMediumSmallSquare           = "◽️"
	EmojiBlackMediumSquare                = "◼️"
	EmojiWhiteMediumSquare                = "◻️"
	EmojiBlackLargeSquare                 = "⬛️"
	EmojiWhiteLargeSquare                 = "⬜️"
	EmojiSpeaker                          = "🔈"
	EmojiMute                             = "🔇"
	EmojiSound                            = "🔉"
	EmojiLoudSound                        = "🔊"
	EmojiBell                             = "🔔"
	EmojiNoBell                           = "🔕"
	EmojiMega                             = "📣"
	EmojiLoudspeaker                      = "📢"
	EmojiEyeSpeechBubble                  = "👁‍🗨"
	EmojiSpeechBalloon                    = "💬"
	EmojiThoughtBalloon                   = "💭"
	EmojiRightAngerBubble                 = "🗯"
	EmojiSpades                           = "♠️"
	EmojiClubs                            = "♣️"
	EmojiHearts                           = "♥️"
	EmojiDiamonds                         = "♦️"
	EmojiBlackJoker                       = "🃏"
	EmojiFlowerPlayingCards               = "🎴"
	EmojiMahjong                          = "🀄️"
	EmojiClock1                           = "🕐"
	EmojiClock2                           = "🕑"
	EmojiClock3                           = "🕒"
	EmojiClock4                           = "🕓"
	EmojiClock5                           = "🕔"
	EmojiClock6                           = "🕕"
	EmojiClock7                           = "🕖"
	EmojiClock8                           = "🕗"
	EmojiClock9                           = "🕘"
	EmojiClock10                          = "🕙"
	EmojiClock11                          = "🕚"
	EmojiClock12                          = "🕛"
	EmojiClock130                         = "🕜"
	EmojiClock230                         = "🕝"
	EmojiClock330                         = "🕞"
	EmojiClock430                         = "🕟"
	EmojiClock530                         = "🕠"
	EmojiClock630                         = "🕡"
	EmojiClock730                         = "🕢"
	EmojiClock830                         = "🕣"
	EmojiClock930                         = "🕤"
	EmojiClock1030                        = "🕥"
	EmojiClock1130                        = "🕦"
	EmojiClock1230                        = "🕧"
	EmojiWhiteFlag                        = "🏳️"
	EmojiBlackFlag                        = "🏴"
	EmojiCheckeredFlag                    = "🏁"
	EmojiTriangularFlagOnPost             = "🚩"
	EmojiRainbowFlag                      = "🏳️‍🌈"
	EmojiAfghanistan                      = "🇦🇫"
	EmojiAlandIslands                     = "🇦🇽"
	EmojiAlbania                          = "🇦🇱"
	EmojiAlgeria                          = "🇩🇿"
	EmojiAmericanSamoa                    = "🇦🇸"
	EmojiAndorra                          = "🇦🇩"
	EmojiAngola                           = "🇦🇴"
	EmojiAnguilla                         = "🇦🇮"
	EmojiAntarctica                       = "🇦🇶"
	EmojiAntiguaBarbuda                   = "🇦🇬"
	EmojiArgentina                        = "🇦🇷"
	EmojiArmenia                          = "🇦🇲"
	EmojiAruba                            = "🇦🇼"
	EmojiAustralia                        = "🇦🇺"
	EmojiAustria                          = "🇦🇹"
	EmojiAzerbaijan                       = "🇦🇿"
	EmojiBahamas                          = "🇧🇸"
	EmojiBahrain                          = "🇧🇭"
	EmojiBangladesh                       = "🇧🇩"
	EmojiBarbados                         = "🇧🇧"
	EmojiBelarus                          = "🇧🇾"
	EmojiBelgium                          = "🇧🇪"
	EmojiBelize                           = "🇧🇿"
	EmojiBenin                            = "🇧🇯"
	EmojiBermuda                          = "🇧🇲"
	EmojiBhutan                           = "🇧🇹"
	EmojiBolivia                          = "🇧🇴"
	EmojiCaribbeanNetherlands             = "🇧🇶"
	EmojiBosniaHerzegovina                = "🇧🇦"
	EmojiBotswana                         = "🇧🇼"
	EmojiBrazil                           = "🇧🇷"
	EmojiBritishIndianOceanTerritory      = "🇮🇴"
	EmojiBritishVirginIslands             = "🇻🇬"
	EmojiBrunei                           = "🇧🇳"
	EmojiBulgaria                         = "🇧🇬"
	EmojiBurkinaFaso                      = "🇧🇫"
	EmojiBurundi                          = "🇧🇮"
	EmojiCapeVerde                        = "🇨🇻"
	EmojiCambodia                         = "🇰🇭"
	EmojiCameroon                         = "🇨🇲"
	EmojiCanada                           = "🇨🇦"
	EmojiCanaryIslands                    = "🇮🇨"
	EmojiCaymanIslands                    = "🇰🇾"
	EmojiCentralAfricanRepublic           = "🇨🇫"
	EmojiChad                             = "🇹🇩"
	EmojiChile                            = "🇨🇱"
	EmojiCn                               = "🇨🇳"
	EmojiChristmasIsland                  = "🇨🇽"
	EmojiCocosIslands                     = "🇨🇨"
	EmojiColombia                         = "🇨🇴"
	EmojiComoros                          = "🇰🇲"
	EmojiCongoBrazzaville                 = "🇨🇬"
	EmojiCongoKinshasa                    = "🇨🇩"
	EmojiCookIslands                      = "🇨🇰"
	EmojiCostaRica                        = "🇨🇷"
	EmojiCoteDivoire                      = "🇨🇮"
	EmojiCroatia                          = "🇭🇷"
	EmojiCuba                             = "🇨🇺"
	EmojiCuracao                          = "🇨🇼"
	EmojiCyprus                           = "🇨🇾"
	EmojiCzechRepublic                    = "🇨🇿"
	EmojiDenmark                          = "🇩🇰"
	EmojiDjibouti                         = "🇩🇯"
	EmojiDominica                         = "🇩🇲"
	EmojiDominicanRepublic                = "🇩🇴"
	EmojiEcuador                          = "🇪🇨"
	EmojiEgypt                            = "🇪🇬"
	EmojiElSalvador                       = "🇸🇻"
	EmojiEquatorialGuinea                 = "🇬🇶"
	EmojiEritrea                          = "🇪🇷"
	EmojiEstonia                          = "🇪🇪"
	EmojiEthiopia                         = "🇪🇹"
	EmojiEu                               = "🇪🇺"
	EmojiFalklandIslands                  = "🇫🇰"
	EmojiFaroeIslands                     = "🇫🇴"
	EmojiFiji                             = "🇫🇯"
	EmojiFinland                          = "🇫🇮"
	EmojiFr                               = "🇫🇷"
	EmojiFrenchGuiana                     = "🇬🇫"
	EmojiFrenchPolynesia                  = "🇵🇫"
	EmojiFrenchSouthernTerritories        = "🇹🇫"
	EmojiGabon                            = "🇬🇦"
	EmojiGambia                           = "🇬🇲"
	EmojiGeorgia                          = "🇬🇪"
	EmojiDe                               = "🇩🇪"
	EmojiGhana                            = "🇬🇭"
	EmojiGibraltar                        = "🇬🇮"
	EmojiGreece                           = "🇬🇷"
	EmojiGreenland                        = "🇬🇱"
	EmojiGrenada                          = "🇬🇩"
	EmojiGuadeloupe                       = "🇬🇵"
	EmojiGuam                             = "🇬🇺"
	EmojiGuatemala                        = "🇬🇹"
	EmojiGuernsey                         = "🇬🇬"
	EmojiGuinea                           = "🇬🇳"
	EmojiGuineaBissau                     = "🇬🇼"
	EmojiGuyana                           = "🇬🇾"
	EmojiHaiti                            = "🇭🇹"
	EmojiHonduras                         = "🇭🇳"
	EmojiHongKong                         = "🇭🇰"
	EmojiHungary                          = "🇭🇺"
	EmojiIceland                          = "🇮🇸"
	EmojiIndia                            = "🇮🇳"
	EmojiIndonesia                        = "🇮🇩"
	EmojiIran                             = "🇮🇷"
	EmojiIraq                             = "🇮🇶"
	EmojiIreland                          = "🇮🇪"
	EmojiIsleOfMan                        = "🇮🇲"
	EmojiIsrael                           = "🇮🇱"
	EmojiIt                               = "🇮🇹"
	EmojiJamaica                          = "🇯🇲"
	EmojiJp                               = "🇯🇵"
	EmojiCrossedFlags                     = "🎌"
	EmojiJersey                           = "🇯🇪"
	EmojiJordan                           = "🇯🇴"
	EmojiKazakhstan                       = "🇰🇿"
	EmojiKenya                            = "🇰🇪"
	EmojiKiribati                         = "🇰🇮"
	EmojiKosovo                           = "🇽🇰"
	EmojiKuwait                           = "🇰🇼"
	EmojiKyrgyzstan                       = "🇰🇬"
	EmojiLaos                             = "🇱🇦"
	EmojiLatvia                           = "🇱🇻"
	EmojiLebanon                          = "🇱🇧"
	EmojiLesotho                          = "🇱🇸"
	EmojiLiberia                          = "🇱🇷"
	EmojiLibya                            = "🇱🇾"
	EmojiLiechtenstein                    = "🇱🇮"
	EmojiLithuania                        = "🇱🇹"
	EmojiLuxembourg                       = "🇱🇺"
	EmojiMacau                            = "🇲🇴"
	EmojiMacedonia                        = "🇲🇰"
	EmojiMadagascar                       = "🇲🇬"
	EmojiMalawi                           = "🇲🇼"
	EmojiMalaysia                         = "🇲🇾"
	EmojiMaldives                         = "🇲🇻"
	EmojiMali                             = "🇲🇱"
	EmojiMalta                            = "🇲🇹"
	EmojiMarshallIslands                  = "🇲🇭"
	EmojiMartinique                       = "🇲🇶"
	EmojiMauritania                       = "🇲🇷"
	EmojiMauritius                        = "🇲🇺"
	EmojiMayotte                          = "🇾🇹"
	EmojiMexico                           = "🇲🇽"
	EmojiMicronesia                       = "🇫🇲"
	EmojiMoldova                          = "🇲🇩"
	EmojiMonaco                           = "🇲🇨"
	EmojiMongolia                         = "🇲🇳"
	EmojiMontenegro                       = "🇲🇪"
	EmojiMontserrat                       = "🇲🇸"
	EmojiMorocco                          = "🇲🇦"
	EmojiMozambique                       = "🇲🇿"
	EmojiMyanmar                          = "🇲🇲"
	EmojiNamibia                          = "🇳🇦"
	EmojiNauru                            = "🇳🇷"
	EmojiNepal                            = "🇳🇵"
	EmojiNetherlands                      = "🇳🇱"
	EmojiNewCaledonia                     = "🇳🇨"
	EmojiNewZealand                       = "🇳🇿"
	EmojiNicaragua                        = "🇳🇮"
	EmojiNiger                            = "🇳🇪"
	EmojiNigeria                          = "🇳🇬"
	EmojiNiue                             = "🇳🇺"
	EmojiNorfolkIsland                    = "🇳🇫"
	EmojiNorthernMarianaIslands           = "🇲🇵"
	EmojiNorthKorea                       = "🇰🇵"
	EmojiNorway                           = "🇳🇴"
	EmojiOman                             = "🇴🇲"
	EmojiPakistan                         = "🇵🇰"
	EmojiPalau                            = "🇵🇼"
	EmojiPalestinianTerritories           = "🇵🇸"
	EmojiPanama                           = "🇵🇦"
	EmojiPapuaNewGuinea                   = "🇵🇬"
	EmojiParaguay                         = "🇵🇾"
	EmojiPeru                             = "🇵🇪"
	EmojiPhilippines                      = "🇵🇭"
	EmojiPitcairnIslands                  = "🇵🇳"
	EmojiPoland                           = "🇵🇱"
	EmojiPortugal                         = "🇵🇹"
	EmojiPuertoRico                       = "🇵🇷"
	EmojiQatar                            = "🇶🇦"
	EmojiReunion                          = "🇷🇪"
	EmojiRomania                          = "🇷🇴"
	EmojiRu                               = "🇷🇺"
	EmojiRwanda                           = "🇷🇼"
	EmojiStBarthelemy                     = "🇧🇱"
	EmojiStHelena                         = "🇸🇭"
	EmojiStKittsNevis                     = "🇰🇳"
	EmojiStLucia                          = "🇱🇨"
	EmojiStPierreMiquelon                 = "🇵🇲"
	EmojiStVincentGrenadines              = "🇻🇨"
	EmojiSamoa                            = "🇼🇸"
	EmojiSanMarino                        = "🇸🇲"
	EmojiSaoTomePrincipe                  = "🇸🇹"
	EmojiSaudiArabia                      = "🇸🇦"
	EmojiSenegal                          = "🇸🇳"
	EmojiSerbia                           = "🇷🇸"
	EmojiSeychelles                       = "🇸🇨"
	EmojiSierraLeone                      = "🇸🇱"
	EmojiSingapore                        = "🇸🇬"
	EmojiSintMaarten                      = "🇸🇽"
	EmojiSlovakia                         = "🇸🇰"
	EmojiSlovenia                         = "🇸🇮"
	EmojiSolomonIslands                   = "🇸🇧"
	EmojiSomalia                          = "🇸🇴"
	EmojiSouthAfrica                      = "🇿🇦"
	EmojiSouthGeorgiaSouthSandwichIslands = "🇬🇸"
	EmojiKr                               = "🇰🇷"
	EmojiSouthSudan                       = "🇸🇸"
	EmojiEs                               = "🇪🇸"
	EmojiSriLanka                         = "🇱🇰"
	EmojiSudan                            = "🇸🇩"
	EmojiSuriname                         = "🇸🇷"
	EmojiSwaziland                        = "🇸🇿"
	EmojiSweden                           = "🇸🇪"
	EmojiSwitzerland                      = "🇨🇭"
	EmojiSyria                            = "🇸🇾"
	EmojiTaiwan                           = "🇹🇼"
	EmojiTajikistan                       = "🇹🇯"
	EmojiTanzania                         = "🇹🇿"
	EmojiThailand                         = "🇹🇭"
	EmojiTimorLeste                       = "🇹🇱"
	EmojiTogo                             = "🇹🇬"
	EmojiTokelau                          = "🇹🇰"
	EmojiTonga                            = "🇹🇴"
	EmojiTrinidadTobago                   = "🇹🇹"
	EmojiTunisia                          = "🇹🇳"
	EmojiTr                               = "🇹🇷"
	EmojiTurkmenistan                     = "🇹🇲"
	EmojiTurksCaicosIslands               = "🇹🇨"
	EmojiTuvalu                           = "🇹🇻"
	EmojiUganda                           = "🇺🇬"
	EmojiUkraine                          = "🇺🇦"
	EmojiUnitedArabEmirates               = "🇦🇪"
	EmojiGb                               = "🇬🇧"
	EmojiUs                               = "🇺🇸"
	EmojiUsVirginIslands                  = "🇻🇮"
	EmojiUruguay                          = "🇺🇾"
	EmojiUzbekistan                       = "🇺🇿"
	EmojiVanuatu                          = "🇻🇺"
	EmojiVaticanCity                      = "🇻🇦"
	EmojiVenezuela                        = "🇻🇪"
	EmojiVietnam                          = "🇻🇳"
	EmojiWallisFutuna                     = "🇼🇫"
	EmojiWesternSahara                    = "🇪🇭"
	EmojiYemen                            = "🇾🇪"
	EmojiZambia                           = "🇿🇲"
	EmojiZimbabwe                         = "🇿🇼"
	EmojiBasecamp                         = ""
	EmojiBasecampy                        = ""
	EmojiBowtie                           = ""
	EmojiFeelsgood                        = ""
	EmojiFinnadie                         = ""
	EmojiGoberserk                        = ""
	EmojiGodmode                          = ""
	EmojiHurtrealbad                      = ""
	EmojiNeckbeard                        = ""
	EmojiOctocat                          = ""
	EmojiRage1                            = ""
	EmojiRage2                            = ""
	EmojiRage3                            = ""
	EmojiRage4                            = ""
	EmojiShipit                           = ""
	EmojiSuspect                          = ""
	EmojiTrollface                        = ""
)

// Emoji returns an emoji from multiple possible aliases
func Emoji(name string) string {
	switch strings.ToLower(name) {

	case "grinning":
		return EmojiGrinning
	case "smiley":
		return EmojiSmiley
	case "smile":
		return EmojiSmile
	case "grin":
		return EmojiGrin
	case "laughing", "satisfied":
		return EmojiLaughing
	case "sweat_smile":
		return EmojiSweatSmile
	case "joy":
		return EmojiJoy
	case "rofl":
		return EmojiRofl
	case "relaxed":
		return EmojiRelaxed
	case "blush":
		return EmojiBlush
	case "innocent":
		return EmojiInnocent
	case "slightly_smiling_face":
		return EmojiSlightlySmilingFace
	case "upside_down_face":
		return EmojiUpsideDownFace
	case "wink":
		return EmojiWink
	case "relieved":
		return EmojiRelieved
	case "heart_eyes":
		return EmojiHeartEyes
	case "kissing_heart":
		return EmojiKissingHeart
	case "kissing":
		return EmojiKissing
	case "kissing_smiling_eyes":
		return EmojiKissingSmilingEyes
	case "kissing_closed_eyes":
		return EmojiKissingClosedEyes
	case "yum":
		return EmojiYum
	case "stuck_out_tongue_winking_eye":
		return EmojiStuckOutTongueWinkingEye
	case "stuck_out_tongue_closed_eyes":
		return EmojiStuckOutTongueClosedEyes
	case "stuck_out_tongue":
		return EmojiStuckOutTongue
	case "money_mouth_face":
		return EmojiMoneyMouthFace
	case "hugs":
		return EmojiHugs
	case "nerd_face":
		return EmojiNerdFace
	case "sunglasses":
		return EmojiSunglasses
	case "clown_face":
		return EmojiClownFace
	case "cowboy_hat_face":
		return EmojiCowboyHatFace
	case "smirk":
		return EmojiSmirk
	case "unamused":
		return EmojiUnamused
	case "disappointed":
		return EmojiDisappointed
	case "pensive":
		return EmojiPensive
	case "worried":
		return EmojiWorried
	case "confused":
		return EmojiConfused
	case "slightly_frowning_face":
		return EmojiSlightlyFrowningFace
	case "frowning_face":
		return EmojiFrowningFace
	case "persevere":
		return EmojiPersevere
	case "confounded":
		return EmojiConfounded
	case "tired_face":
		return EmojiTiredFace
	case "weary":
		return EmojiWeary
	case "triumph":
		return EmojiTriumph
	case "angry":
		return EmojiAngry
	case "rage", "pout":
		return EmojiRage
	case "no_mouth":
		return EmojiNoMouth
	case "neutral_face":
		return EmojiNeutralFace
	case "expressionless":
		return EmojiExpressionless
	case "hushed":
		return EmojiHushed
	case "frowning":
		return EmojiFrowning
	case "anguished":
		return EmojiAnguished
	case "open_mouth":
		return EmojiOpenMouth
	case "astonished":
		return EmojiAstonished
	case "dizzy_face":
		return EmojiDizzyFace
	case "flushed":
		return EmojiFlushed
	case "scream":
		return EmojiScream
	case "fearful":
		return EmojiFearful
	case "cold_sweat":
		return EmojiColdSweat
	case "cry":
		return EmojiCry
	case "disappointed_relieved":
		return EmojiDisappointedRelieved
	case "drooling_face":
		return EmojiDroolingFace
	case "sob":
		return EmojiSob
	case "sweat":
		return EmojiSweat
	case "sleepy":
		return EmojiSleepy
	case "sleeping":
		return EmojiSleeping
	case "roll_eyes":
		return EmojiRollEyes
	case "thinking":
		return EmojiThinking
	case "lying_face":
		return EmojiLyingFace
	case "grimacing":
		return EmojiGrimacing
	case "zipper_mouth_face":
		return EmojiZipperMouthFace
	case "nauseated_face":
		return EmojiNauseatedFace
	case "sneezing_face":
		return EmojiSneezingFace
	case "mask":
		return EmojiMask
	case "face_with_thermometer":
		return EmojiFaceWithThermometer
	case "face_with_head_bandage":
		return EmojiFaceWithHeadBandage
	case "smiling_imp":
		return EmojiSmilingImp
	case "imp":
		return EmojiImp
	case "japanese_ogre":
		return EmojiJapaneseOgre
	case "japanese_goblin":
		return EmojiJapaneseGoblin
	case "hankey", "poop", "shit":
		return EmojiHankey
	case "ghost":
		return EmojiGhost
	case "skull":
		return EmojiSkull
	case "skull_and_crossbones":
		return EmojiSkullAndCrossbones
	case "alien":
		return EmojiAlien
	case "space_invader":
		return EmojiSpaceInvader
	case "robot":
		return EmojiRobot
	case "jack_o_lantern":
		return EmojiJackOLantern
	case "smiley_cat":
		return EmojiSmileyCat
	case "smile_cat":
		return EmojiSmileCat
	case "joy_cat":
		return EmojiJoyCat
	case "heart_eyes_cat":
		return EmojiHeartEyesCat
	case "smirk_cat":
		return EmojiSmirkCat
	case "kissing_cat":
		return EmojiKissingCat
	case "scream_cat":
		return EmojiScreamCat
	case "crying_cat_face":
		return EmojiCryingCatFace
	case "pouting_cat":
		return EmojiPoutingCat
	case "open_hands":
		return EmojiOpenHands
	case "raised_hands":
		return EmojiRaisedHands
	case "clap":
		return EmojiClap
	case "pray":
		return EmojiPray
	case "handshake":
		return EmojiHandshake
	case "+1", "thumbsup":
		return EmojiThumbsup
	case "-1", "thumbsdown":
		return EmojiThumbsdown
	case "fist_oncoming", "facepunch", "punch":
		return EmojiFistOncoming
	case "fist_raised", "fist":
		return EmojiFistRaised
	case "fist_left":
		return EmojiFistLeft
	case "fist_right":
		return EmojiFistRight
	case "crossed_fingers":
		return EmojiCrossedFingers
	case "v":
		return EmojiV
	case "metal":
		return EmojiMetal
	case "ok_hand":
		return EmojiOkHand
	case "point_left":
		return EmojiPointLeft
	case "point_right":
		return EmojiPointRight
	case "point_up_2":
		return EmojiPointUp2
	case "point_down":
		return EmojiPointDown
	case "point_up":
		return EmojiPointUp
	case "hand", "raised_hand":
		return EmojiHand
	case "raised_back_of_hand":
		return EmojiRaisedBackOfHand
	case "raised_hand_with_fingers_splayed":
		return EmojiRaisedHandWithFingersSplayed
	case "vulcan_salute":
		return EmojiVulcanSalute
	case "wave":
		return EmojiWave
	case "call_me_hand":
		return EmojiCallMeHand
	case "muscle":
		return EmojiMuscle
	case "middle_finger", "fu":
		return EmojiMiddleFinger
	case "writing_hand":
		return EmojiWritingHand
	case "selfie":
		return EmojiSelfie
	case "nail_care":
		return EmojiNailCare
	case "ring":
		return EmojiRing
	case "lipstick":
		return EmojiLipstick
	case "kiss":
		return EmojiKiss
	case "lips":
		return EmojiLips
	case "tongue":
		return EmojiTongue
	case "ear":
		return EmojiEar
	case "nose":
		return EmojiNose
	case "footprints":
		return EmojiFootprints
	case "eye":
		return EmojiEye
	case "eyes":
		return EmojiEyes
	case "speaking_head":
		return EmojiSpeakingHead
	case "bust_in_silhouette":
		return EmojiBustInSilhouette
	case "busts_in_silhouette":
		return EmojiBustsInSilhouette
	case "baby":
		return EmojiBaby
	case "boy":
		return EmojiBoy
	case "girl":
		return EmojiGirl
	case "man":
		return EmojiMan
	case "woman":
		return EmojiWoman
	case "blonde_woman":
		return EmojiBlondeWoman
	case "blonde_man", "person_with_blond_hair":
		return EmojiBlondeMan
	case "older_man":
		return EmojiOlderMan
	case "older_woman":
		return EmojiOlderWoman
	case "man_with_gua_pi_mao":
		return EmojiManWithGuaPiMao
	case "woman_with_turban":
		return EmojiWomanWithTurban
	case "man_with_turban":
		return EmojiManWithTurban
	case "policewoman":
		return EmojiPolicewoman
	case "policeman", "cop":
		return EmojiPoliceman
	case "construction_worker_woman":
		return EmojiConstructionWorkerWoman
	case "construction_worker_man", "construction_worker":
		return EmojiConstructionWorkerMan
	case "guardswoman":
		return EmojiGuardswoman
	case "guardsman":
		return EmojiGuardsman
	case "female_detective":
		return EmojiFemaleDetective
	case "male_detective", "detective":
		return EmojiMaleDetective
	case "woman_health_worker":
		return EmojiWomanHealthWorker
	case "man_health_worker":
		return EmojiManHealthWorker
	case "woman_farmer":
		return EmojiWomanFarmer
	case "man_farmer":
		return EmojiManFarmer
	case "woman_cook":
		return EmojiWomanCook
	case "man_cook":
		return EmojiManCook
	case "woman_student":
		return EmojiWomanStudent
	case "man_student":
		return EmojiManStudent
	case "woman_singer":
		return EmojiWomanSinger
	case "man_singer":
		return EmojiManSinger
	case "woman_teacher":
		return EmojiWomanTeacher
	case "man_teacher":
		return EmojiManTeacher
	case "woman_factory_worker":
		return EmojiWomanFactoryWorker
	case "man_factory_worker":
		return EmojiManFactoryWorker
	case "woman_technologist":
		return EmojiWomanTechnologist
	case "man_technologist":
		return EmojiManTechnologist
	case "woman_office_worker":
		return EmojiWomanOfficeWorker
	case "man_office_worker":
		return EmojiManOfficeWorker
	case "woman_mechanic":
		return EmojiWomanMechanic
	case "man_mechanic":
		return EmojiManMechanic
	case "woman_scientist":
		return EmojiWomanScientist
	case "man_scientist":
		return EmojiManScientist
	case "woman_artist":
		return EmojiWomanArtist
	case "man_artist":
		return EmojiManArtist
	case "woman_firefighter":
		return EmojiWomanFirefighter
	case "man_firefighter":
		return EmojiManFirefighter
	case "woman_pilot":
		return EmojiWomanPilot
	case "man_pilot":
		return EmojiManPilot
	case "woman_astronaut":
		return EmojiWomanAstronaut
	case "man_astronaut":
		return EmojiManAstronaut
	case "woman_judge":
		return EmojiWomanJudge
	case "man_judge":
		return EmojiManJudge
	case "mrs_claus":
		return EmojiMrsClaus
	case "santa":
		return EmojiSanta
	case "princess":
		return EmojiPrincess
	case "prince":
		return EmojiPrince
	case "bride_with_veil":
		return EmojiBrideWithVeil
	case "man_in_tuxedo":
		return EmojiManInTuxedo
	case "angel":
		return EmojiAngel
	case "pregnant_woman":
		return EmojiPregnantWoman
	case "bowing_woman":
		return EmojiBowingWoman
	case "bowing_man", "bow":
		return EmojiBowingMan
	case "tipping_hand_woman", "information_desk_person", "sassy_woman":
		return EmojiTippingHandWoman
	case "tipping_hand_man", "sassy_man":
		return EmojiTippingHandMan
	case "no_good_woman", "no_good", "ng_woman":
		return EmojiNoGoodWoman
	case "no_good_man", "ng_man":
		return EmojiNoGoodMan
	case "ok_woman":
		return EmojiOkWoman
	case "ok_man":
		return EmojiOkMan
	case "raising_hand_woman", "raising_hand":
		return EmojiRaisingHandWoman
	case "raising_hand_man":
		return EmojiRaisingHandMan
	case "woman_facepalming":
		return EmojiWomanFacepalming
	case "man_facepalming":
		return EmojiManFacepalming
	case "woman_shrugging":
		return EmojiWomanShrugging
	case "man_shrugging":
		return EmojiManShrugging
	case "pouting_woman", "person_with_pouting_face":
		return EmojiPoutingWoman
	case "pouting_man":
		return EmojiPoutingMan
	case "frowning_woman", "person_frowning":
		return EmojiFrowningWoman
	case "frowning_man":
		return EmojiFrowningMan
	case "haircut_woman", "haircut":
		return EmojiHaircutWoman
	case "haircut_man":
		return EmojiHaircutMan
	case "massage_woman", "massage":
		return EmojiMassageWoman
	case "massage_man":
		return EmojiMassageMan
	case "business_suit_levitating":
		return EmojiBusinessSuitLevitating
	case "dancer":
		return EmojiDancer
	case "man_dancing":
		return EmojiManDancing
	case "dancing_women", "dancers":
		return EmojiDancingWomen
	case "dancing_men":
		return EmojiDancingMen
	case "walking_woman":
		return EmojiWalkingWoman
	case "walking_man", "walking":
		return EmojiWalkingMan
	case "running_woman":
		return EmojiRunningWoman
	case "running_man", "runner", "running":
		return EmojiRunningMan
	case "couple":
		return EmojiCouple
	case "two_women_holding_hands":
		return EmojiTwoWomenHoldingHands
	case "two_men_holding_hands":
		return EmojiTwoMenHoldingHands
	case "couple_with_heart_woman_man", "couple_with_heart":
		return EmojiCoupleWithHeartWomanMan
	case "couple_with_heart_woman_woman":
		return EmojiCoupleWithHeartWomanWoman
	case "couple_with_heart_man_man":
		return EmojiCoupleWithHeartManMan
	case "couplekiss_man_woman":
		return EmojiCouplekissManWoman
	case "couplekiss_woman_woman":
		return EmojiCouplekissWomanWoman
	case "couplekiss_man_man":
		return EmojiCouplekissManMan
	case "family_man_woman_boy", "family":
		return EmojiFamilyManWomanBoy
	case "family_man_woman_girl":
		return EmojiFamilyManWomanGirl
	case "family_man_woman_girl_boy":
		return EmojiFamilyManWomanGirlBoy
	case "family_man_woman_boy_boy":
		return EmojiFamilyManWomanBoyBoy
	case "family_man_woman_girl_girl":
		return EmojiFamilyManWomanGirlGirl
	case "family_woman_woman_boy":
		return EmojiFamilyWomanWomanBoy
	case "family_woman_woman_girl":
		return EmojiFamilyWomanWomanGirl
	case "family_woman_woman_girl_boy":
		return EmojiFamilyWomanWomanGirlBoy
	case "family_woman_woman_boy_boy":
		return EmojiFamilyWomanWomanBoyBoy
	case "family_woman_woman_girl_girl":
		return EmojiFamilyWomanWomanGirlGirl
	case "family_man_man_boy":
		return EmojiFamilyManManBoy
	case "family_man_man_girl":
		return EmojiFamilyManManGirl
	case "family_man_man_girl_boy":
		return EmojiFamilyManManGirlBoy
	case "family_man_man_boy_boy":
		return EmojiFamilyManManBoyBoy
	case "family_man_man_girl_girl":
		return EmojiFamilyManManGirlGirl
	case "family_woman_boy":
		return EmojiFamilyWomanBoy
	case "family_woman_girl":
		return EmojiFamilyWomanGirl
	case "family_woman_girl_boy":
		return EmojiFamilyWomanGirlBoy
	case "family_woman_boy_boy":
		return EmojiFamilyWomanBoyBoy
	case "family_woman_girl_girl":
		return EmojiFamilyWomanGirlGirl
	case "family_man_boy":
		return EmojiFamilyManBoy
	case "family_man_girl":
		return EmojiFamilyManGirl
	case "family_man_girl_boy":
		return EmojiFamilyManGirlBoy
	case "family_man_boy_boy":
		return EmojiFamilyManBoyBoy
	case "family_man_girl_girl":
		return EmojiFamilyManGirlGirl
	case "womans_clothes":
		return EmojiWomansClothes
	case "shirt", "tshirt":
		return EmojiShirt
	case "jeans":
		return EmojiJeans
	case "necktie":
		return EmojiNecktie
	case "dress":
		return EmojiDress
	case "bikini":
		return EmojiBikini
	case "kimono":
		return EmojiKimono
	case "high_heel":
		return EmojiHighHeel
	case "sandal":
		return EmojiSandal
	case "boot":
		return EmojiBoot
	case "mans_shoe", "shoe":
		return EmojiMansShoe
	case "athletic_shoe":
		return EmojiAthleticShoe
	case "womans_hat":
		return EmojiWomansHat
	case "tophat":
		return EmojiTophat
	case "mortar_board":
		return EmojiMortarBoard
	case "crown":
		return EmojiCrown
	case "rescue_worker_helmet":
		return EmojiRescueWorkerHelmet
	case "school_satchel":
		return EmojiSchoolSatchel
	case "pouch":
		return EmojiPouch
	case "purse":
		return EmojiPurse
	case "handbag":
		return EmojiHandbag
	case "briefcase":
		return EmojiBriefcase
	case "eyeglasses":
		return EmojiEyeglasses
	case "dark_sunglasses":
		return EmojiDarkSunglasses
	case "closed_umbrella":
		return EmojiClosedUmbrella
	case "open_umbrella":
		return EmojiOpenUmbrella
	case "dog":
		return EmojiDog
	case "cat":
		return EmojiCat
	case "mouse":
		return EmojiMouse
	case "hamster":
		return EmojiHamster
	case "rabbit":
		return EmojiRabbit
	case "fox_face":
		return EmojiFoxFace
	case "bear":
		return EmojiBear
	case "panda_face":
		return EmojiPandaFace
	case "koala":
		return EmojiKoala
	case "tiger":
		return EmojiTiger
	case "lion":
		return EmojiLion
	case "cow":
		return EmojiCow
	case "pig":
		return EmojiPig
	case "pig_nose":
		return EmojiPigNose
	case "frog":
		return EmojiFrog
	case "monkey_face":
		return EmojiMonkeyFace
	case "see_no_evil":
		return EmojiSeeNoEvil
	case "hear_no_evil":
		return EmojiHearNoEvil
	case "speak_no_evil":
		return EmojiSpeakNoEvil
	case "monkey":
		return EmojiMonkey
	case "chicken":
		return EmojiChicken
	case "penguin":
		return EmojiPenguin
	case "bird":
		return EmojiBird
	case "baby_chick":
		return EmojiBabyChick
	case "hatching_chick":
		return EmojiHatchingChick
	case "hatched_chick":
		return EmojiHatchedChick
	case "duck":
		return EmojiDuck
	case "eagle":
		return EmojiEagle
	case "owl":
		return EmojiOwl
	case "bat":
		return EmojiBat
	case "wolf":
		return EmojiWolf
	case "boar":
		return EmojiBoar
	case "horse":
		return EmojiHorse
	case "unicorn":
		return EmojiUnicorn
	case "bee", "honeybee":
		return EmojiBee
	case "bug":
		return EmojiBug
	case "butterfly":
		return EmojiButterfly
	case "snail":
		return EmojiSnail
	case "shell":
		return EmojiShell
	case "beetle":
		return EmojiBeetle
	case "ant":
		return EmojiAnt
	case "spider":
		return EmojiSpider
	case "spider_web":
		return EmojiSpiderWeb
	case "turtle":
		return EmojiTurtle
	case "snake":
		return EmojiSnake
	case "lizard":
		return EmojiLizard
	case "scorpion":
		return EmojiScorpion
	case "crab":
		return EmojiCrab
	case "squid":
		return EmojiSquid
	case "octopus":
		return EmojiOctopus
	case "shrimp":
		return EmojiShrimp
	case "tropical_fish":
		return EmojiTropicalFish
	case "fish":
		return EmojiFish
	case "blowfish":
		return EmojiBlowfish
	case "dolphin", "flipper":
		return EmojiDolphin
	case "shark":
		return EmojiShark
	case "whale":
		return EmojiWhale
	case "whale2":
		return EmojiWhale2
	case "crocodile":
		return EmojiCrocodile
	case "leopard":
		return EmojiLeopard
	case "tiger2":
		return EmojiTiger2
	case "water_buffalo":
		return EmojiWaterBuffalo
	case "ox":
		return EmojiOx
	case "cow2":
		return EmojiCow2
	case "deer":
		return EmojiDeer
	case "dromedary_camel":
		return EmojiDromedaryCamel
	case "camel":
		return EmojiCamel
	case "elephant":
		return EmojiElephant
	case "rhinoceros":
		return EmojiRhinoceros
	case "gorilla":
		return EmojiGorilla
	case "racehorse":
		return EmojiRacehorse
	case "pig2":
		return EmojiPig2
	case "goat":
		return EmojiGoat
	case "ram":
		return EmojiRam
	case "sheep":
		return EmojiSheep
	case "dog2":
		return EmojiDog2
	case "poodle":
		return EmojiPoodle
	case "cat2":
		return EmojiCat2
	case "rooster":
		return EmojiRooster
	case "turkey":
		return EmojiTurkey
	case "dove":
		return EmojiDove
	case "rabbit2":
		return EmojiRabbit2
	case "mouse2":
		return EmojiMouse2
	case "rat":
		return EmojiRat
	case "chipmunk":
		return EmojiChipmunk
	case "feet", "paw_prints":
		return EmojiFeet
	case "dragon":
		return EmojiDragon
	case "dragon_face":
		return EmojiDragonFace
	case "cactus":
		return EmojiCactus
	case "christmas_tree":
		return EmojiChristmasTree
	case "evergreen_tree":
		return EmojiEvergreenTree
	case "deciduous_tree":
		return EmojiDeciduousTree
	case "palm_tree":
		return EmojiPalmTree
	case "seedling":
		return EmojiSeedling
	case "herb":
		return EmojiHerb
	case "shamrock":
		return EmojiShamrock
	case "four_leaf_clover":
		return EmojiFourLeafClover
	case "bamboo":
		return EmojiBamboo
	case "tanabata_tree":
		return EmojiTanabataTree
	case "leaves":
		return EmojiLeaves
	case "fallen_leaf":
		return EmojiFallenLeaf
	case "maple_leaf":
		return EmojiMapleLeaf
	case "mushroom":
		return EmojiMushroom
	case "ear_of_rice":
		return EmojiEarOfRice
	case "bouquet":
		return EmojiBouquet
	case "tulip":
		return EmojiTulip
	case "rose":
		return EmojiRose
	case "wilted_flower":
		return EmojiWiltedFlower
	case "sunflower":
		return EmojiSunflower
	case "blossom":
		return EmojiBlossom
	case "cherry_blossom":
		return EmojiCherryBlossom
	case "hibiscus":
		return EmojiHibiscus
	case "earth_americas":
		return EmojiEarthAmericas
	case "earth_africa":
		return EmojiEarthAfrica
	case "earth_asia":
		return EmojiEarthAsia
	case "full_moon":
		return EmojiFullMoon
	case "waning_gibbous_moon":
		return EmojiWaningGibbousMoon
	case "last_quarter_moon":
		return EmojiLastQuarterMoon
	case "waning_crescent_moon":
		return EmojiWaningCrescentMoon
	case "new_moon":
		return EmojiNewMoon
	case "waxing_crescent_moon":
		return EmojiWaxingCrescentMoon
	case "first_quarter_moon":
		return EmojiFirstQuarterMoon
	case "moon", "waxing_gibbous_moon":
		return EmojiMoon
	case "new_moon_with_face":
		return EmojiNewMoonWithFace
	case "full_moon_with_face":
		return EmojiFullMoonWithFace
	case "sun_with_face":
		return EmojiSunWithFace
	case "first_quarter_moon_with_face":
		return EmojiFirstQuarterMoonWithFace
	case "last_quarter_moon_with_face":
		return EmojiLastQuarterMoonWithFace
	case "crescent_moon":
		return EmojiCrescentMoon
	case "dizzy":
		return EmojiDizzy
	case "star":
		return EmojiStar
	case "star2":
		return EmojiStar2
	case "sparkles":
		return EmojiSparkles
	case "zap":
		return EmojiZap
	case "fire":
		return EmojiFire
	case "boom", "collision":
		return EmojiBoom
	case "comet":
		return EmojiComet
	case "sunny":
		return EmojiSunny
	case "sun_behind_small_cloud":
		return EmojiSunBehindSmallCloud
	case "partly_sunny":
		return EmojiPartlySunny
	case "sun_behind_large_cloud":
		return EmojiSunBehindLargeCloud
	case "sun_behind_rain_cloud":
		return EmojiSunBehindRainCloud
	case "rainbow":
		return EmojiRainbow
	case "cloud":
		return EmojiCloud
	case "cloud_with_rain":
		return EmojiCloudWithRain
	case "cloud_with_lightning_and_rain":
		return EmojiCloudWithLightningAndRain
	case "cloud_with_lightning":
		return EmojiCloudWithLightning
	case "cloud_with_snow":
		return EmojiCloudWithSnow
	case "snowman_with_snow":
		return EmojiSnowmanWithSnow
	case "snowman":
		return EmojiSnowman
	case "snowflake":
		return EmojiSnowflake
	case "wind_face":
		return EmojiWindFace
	case "dash":
		return EmojiDash
	case "tornado":
		return EmojiTornado
	case "fog":
		return EmojiFog
	case "ocean":
		return EmojiOcean
	case "droplet":
		return EmojiDroplet
	case "sweat_drops":
		return EmojiSweatDrops
	case "umbrella":
		return EmojiUmbrella
	case "green_apple":
		return EmojiGreenApple
	case "apple":
		return EmojiApple
	case "pear":
		return EmojiPear
	case "tangerine", "orange", "mandarin":
		return EmojiTangerine
	case "lemon":
		return EmojiLemon
	case "banana":
		return EmojiBanana
	case "watermelon":
		return EmojiWatermelon
	case "grapes":
		return EmojiGrapes
	case "strawberry":
		return EmojiStrawberry
	case "melon":
		return EmojiMelon
	case "cherries":
		return EmojiCherries
	case "peach":
		return EmojiPeach
	case "pineapple":
		return EmojiPineapple
	case "kiwi_fruit":
		return EmojiKiwiFruit
	case "avocado":
		return EmojiAvocado
	case "tomato":
		return EmojiTomato
	case "eggplant":
		return EmojiEggplant
	case "cucumber":
		return EmojiCucumber
	case "carrot":
		return EmojiCarrot
	case "corn":
		return EmojiCorn
	case "hot_pepper":
		return EmojiHotPepper
	case "potato":
		return EmojiPotato
	case "sweet_potato":
		return EmojiSweetPotato
	case "chestnut":
		return EmojiChestnut
	case "peanuts":
		return EmojiPeanuts
	case "honey_pot":
		return EmojiHoneyPot
	case "croissant":
		return EmojiCroissant
	case "bread":
		return EmojiBread
	case "baguette_bread":
		return EmojiBaguetteBread
	case "cheese":
		return EmojiCheese
	case "egg":
		return EmojiEgg
	case "fried_egg":
		return EmojiFriedEgg
	case "bacon":
		return EmojiBacon
	case "pancakes":
		return EmojiPancakes
	case "fried_shrimp":
		return EmojiFriedShrimp
	case "poultry_leg":
		return EmojiPoultryLeg
	case "meat_on_bone":
		return EmojiMeatOnBone
	case "pizza":
		return EmojiPizza
	case "hotdog":
		return EmojiHotdog
	case "hamburger":
		return EmojiHamburger
	case "fries":
		return EmojiFries
	case "stuffed_flatbread":
		return EmojiStuffedFlatbread
	case "taco":
		return EmojiTaco
	case "burrito":
		return EmojiBurrito
	case "green_salad":
		return EmojiGreenSalad
	case "shallow_pan_of_food":
		return EmojiShallowPanOfFood
	case "spaghetti":
		return EmojiSpaghetti
	case "ramen":
		return EmojiRamen
	case "stew":
		return EmojiStew
	case "fish_cake":
		return EmojiFishCake
	case "sushi":
		return EmojiSushi
	case "bento":
		return EmojiBento
	case "curry":
		return EmojiCurry
	case "rice":
		return EmojiRice
	case "rice_ball":
		return EmojiRiceBall
	case "rice_cracker":
		return EmojiRiceCracker
	case "oden":
		return EmojiOden
	case "dango":
		return EmojiDango
	case "shaved_ice":
		return EmojiShavedIce
	case "ice_cream":
		return EmojiIceCream
	case "icecream":
		return EmojiIcecream
	case "cake":
		return EmojiCake
	case "birthday":
		return EmojiBirthday
	case "custard":
		return EmojiCustard
	case "lollipop":
		return EmojiLollipop
	case "candy":
		return EmojiCandy
	case "chocolate_bar":
		return EmojiChocolateBar
	case "popcorn":
		return EmojiPopcorn
	case "doughnut":
		return EmojiDoughnut
	case "cookie":
		return EmojiCookie
	case "milk_glass":
		return EmojiMilkGlass
	case "baby_bottle":
		return EmojiBabyBottle
	case "coffee":
		return EmojiCoffee
	case "tea":
		return EmojiTea
	case "sake":
		return EmojiSake
	case "beer":
		return EmojiBeer
	case "beers":
		return EmojiBeers
	case "clinking_glasses":
		return EmojiClinkingGlasses
	case "wine_glass":
		return EmojiWineGlass
	case "tumbler_glass":
		return EmojiTumblerGlass
	case "cocktail":
		return EmojiCocktail
	case "tropical_drink":
		return EmojiTropicalDrink
	case "champagne":
		return EmojiChampagne
	case "spoon":
		return EmojiSpoon
	case "fork_and_knife":
		return EmojiForkAndKnife
	case "plate_with_cutlery":
		return EmojiPlateWithCutlery
	case "soccer":
		return EmojiSoccer
	case "basketball":
		return EmojiBasketball
	case "football":
		return EmojiFootball
	case "baseball":
		return EmojiBaseball
	case "tennis":
		return EmojiTennis
	case "volleyball":
		return EmojiVolleyball
	case "rugby_football":
		return EmojiRugbyFootball
	case "8ball":
		return Emoji8ball
	case "ping_pong":
		return EmojiPingPong
	case "badminton":
		return EmojiBadminton
	case "goal_net":
		return EmojiGoalNet
	case "ice_hockey":
		return EmojiIceHockey
	case "field_hockey":
		return EmojiFieldHockey
	case "cricket":
		return EmojiCricket
	case "golf":
		return EmojiGolf
	case "bow_and_arrow":
		return EmojiBowAndArrow
	case "fishing_pole_and_fish":
		return EmojiFishingPoleAndFish
	case "boxing_glove":
		return EmojiBoxingGlove
	case "martial_arts_uniform":
		return EmojiMartialArtsUniform
	case "ice_skate":
		return EmojiIceSkate
	case "ski":
		return EmojiSki
	case "skier":
		return EmojiSkier
	case "snowboarder":
		return EmojiSnowboarder
	case "weight_lifting_woman":
		return EmojiWeightLiftingWoman
	case "weight_lifting_man":
		return EmojiWeightLiftingMan
	case "person_fencing":
		return EmojiPersonFencing
	case "women_wrestling":
		return EmojiWomenWrestling
	case "men_wrestling":
		return EmojiMenWrestling
	case "woman_cartwheeling":
		return EmojiWomanCartwheeling
	case "man_cartwheeling":
		return EmojiManCartwheeling
	case "basketball_woman":
		return EmojiBasketballWoman
	case "basketball_man":
		return EmojiBasketballMan
	case "woman_playing_handball":
		return EmojiWomanPlayingHandball
	case "man_playing_handball":
		return EmojiManPlayingHandball
	case "golfing_woman":
		return EmojiGolfingWoman
	case "golfing_man":
		return EmojiGolfingMan
	case "surfing_woman":
		return EmojiSurfingWoman
	case "surfing_man", "surfer":
		return EmojiSurfingMan
	case "swimming_woman":
		return EmojiSwimmingWoman
	case "swimming_man", "swimmer":
		return EmojiSwimmingMan
	case "woman_playing_water_polo":
		return EmojiWomanPlayingWaterPolo
	case "man_playing_water_polo":
		return EmojiManPlayingWaterPolo
	case "rowing_woman":
		return EmojiRowingWoman
	case "rowing_man", "rowboat":
		return EmojiRowingMan
	case "horse_racing":
		return EmojiHorseRacing
	case "biking_woman":
		return EmojiBikingWoman
	case "biking_man", "bicyclist":
		return EmojiBikingMan
	case "mountain_biking_woman":
		return EmojiMountainBikingWoman
	case "mountain_biking_man", "mountain_bicyclist":
		return EmojiMountainBikingMan
	case "running_shirt_with_sash":
		return EmojiRunningShirtWithSash
	case "medal_sports":
		return EmojiMedalSports
	case "medal_military":
		return EmojiMedalMilitary
	case "1st_place_medal":
		return Emoji1stPlaceMedal
	case "2nd_place_medal":
		return Emoji2ndPlaceMedal
	case "3rd_place_medal":
		return Emoji3rdPlaceMedal
	case "trophy":
		return EmojiTrophy
	case "rosette":
		return EmojiRosette
	case "reminder_ribbon":
		return EmojiReminderRibbon
	case "ticket":
		return EmojiTicket
	case "tickets":
		return EmojiTickets
	case "circus_tent":
		return EmojiCircusTent
	case "woman_juggling":
		return EmojiWomanJuggling
	case "man_juggling":
		return EmojiManJuggling
	case "performing_arts":
		return EmojiPerformingArts
	case "art":
		return EmojiArt
	case "clapper":
		return EmojiClapper
	case "microphone":
		return EmojiMicrophone
	case "headphones":
		return EmojiHeadphones
	case "musical_score":
		return EmojiMusicalScore
	case "musical_keyboard":
		return EmojiMusicalKeyboard
	case "drum":
		return EmojiDrum
	case "saxophone":
		return EmojiSaxophone
	case "trumpet":
		return EmojiTrumpet
	case "guitar":
		return EmojiGuitar
	case "violin":
		return EmojiViolin
	case "game_die":
		return EmojiGameDie
	case "dart":
		return EmojiDart
	case "bowling":
		return EmojiBowling
	case "video_game":
		return EmojiVideoGame
	case "slot_machine":
		return EmojiSlotMachine
	case "car", "red_car":
		return EmojiCar
	case "taxi":
		return EmojiTaxi
	case "blue_car":
		return EmojiBlueCar
	case "bus":
		return EmojiBus
	case "trolleybus":
		return EmojiTrolleybus
	case "racing_car":
		return EmojiRacingCar
	case "police_car":
		return EmojiPoliceCar
	case "ambulance":
		return EmojiAmbulance
	case "fire_engine":
		return EmojiFireEngine
	case "minibus":
		return EmojiMinibus
	case "truck":
		return EmojiTruck
	case "articulated_lorry":
		return EmojiArticulatedLorry
	case "tractor":
		return EmojiTractor
	case "kick_scooter":
		return EmojiKickScooter
	case "bike":
		return EmojiBike
	case "motor_scooter":
		return EmojiMotorScooter
	case "motorcycle":
		return EmojiMotorcycle
	case "rotating_light":
		return EmojiRotatingLight
	case "oncoming_police_car":
		return EmojiOncomingPoliceCar
	case "oncoming_bus":
		return EmojiOncomingBus
	case "oncoming_automobile":
		return EmojiOncomingAutomobile
	case "oncoming_taxi":
		return EmojiOncomingTaxi
	case "aerial_tramway":
		return EmojiAerialTramway
	case "mountain_cableway":
		return EmojiMountainCableway
	case "suspension_railway":
		return EmojiSuspensionRailway
	case "railway_car":
		return EmojiRailwayCar
	case "train":
		return EmojiTrain
	case "mountain_railway":
		return EmojiMountainRailway
	case "monorail":
		return EmojiMonorail
	case "bullettrain_side":
		return EmojiBullettrainSide
	case "bullettrain_front":
		return EmojiBullettrainFront
	case "light_rail":
		return EmojiLightRail
	case "steam_locomotive":
		return EmojiSteamLocomotive
	case "train2":
		return EmojiTrain2
	case "metro":
		return EmojiMetro
	case "tram":
		return EmojiTram
	case "station":
		return EmojiStation
	case "helicopter":
		return EmojiHelicopter
	case "small_airplane":
		return EmojiSmallAirplane
	case "airplane":
		return EmojiAirplane
	case "flight_departure":
		return EmojiFlightDeparture
	case "flight_arrival":
		return EmojiFlightArrival
	case "rocket":
		return EmojiRocket
	case "artificial_satellite":
		return EmojiArtificialSatellite
	case "seat":
		return EmojiSeat
	case "canoe":
		return EmojiCanoe
	case "boat", "sailboat":
		return EmojiBoat
	case "motor_boat":
		return EmojiMotorBoat
	case "speedboat":
		return EmojiSpeedboat
	case "passenger_ship":
		return EmojiPassengerShip
	case "ferry":
		return EmojiFerry
	case "ship":
		return EmojiShip
	case "anchor":
		return EmojiAnchor
	case "construction":
		return EmojiConstruction
	case "fuelpump":
		return EmojiFuelpump
	case "busstop":
		return EmojiBusstop
	case "vertical_traffic_light":
		return EmojiVerticalTrafficLight
	case "traffic_light":
		return EmojiTrafficLight
	case "world_map":
		return EmojiWorldMap
	case "moyai":
		return EmojiMoyai
	case "statue_of_liberty":
		return EmojiStatueOfLiberty
	case "fountain":
		return EmojiFountain
	case "tokyo_tower":
		return EmojiTokyoTower
	case "european_castle":
		return EmojiEuropeanCastle
	case "japanese_castle":
		return EmojiJapaneseCastle
	case "stadium":
		return EmojiStadium
	case "ferris_wheel":
		return EmojiFerrisWheel
	case "roller_coaster":
		return EmojiRollerCoaster
	case "carousel_horse":
		return EmojiCarouselHorse
	case "parasol_on_ground":
		return EmojiParasolOnGround
	case "beach_umbrella":
		return EmojiBeachUmbrella
	case "desert_island":
		return EmojiDesertIsland
	case "mountain":
		return EmojiMountain
	case "mountain_snow":
		return EmojiMountainSnow
	case "mount_fuji":
		return EmojiMountFuji
	case "volcano":
		return EmojiVolcano
	case "desert":
		return EmojiDesert
	case "camping":
		return EmojiCamping
	case "tent":
		return EmojiTent
	case "railway_track":
		return EmojiRailwayTrack
	case "motorway":
		return EmojiMotorway
	case "building_construction":
		return EmojiBuildingConstruction
	case "factory":
		return EmojiFactory
	case "house":
		return EmojiHouse
	case "house_with_garden":
		return EmojiHouseWithGarden
	case "houses":
		return EmojiHouses
	case "derelict_house":
		return EmojiDerelictHouse
	case "office":
		return EmojiOffice
	case "department_store":
		return EmojiDepartmentStore
	case "post_office":
		return EmojiPostOffice
	case "european_post_office":
		return EmojiEuropeanPostOffice
	case "hospital":
		return EmojiHospital
	case "bank":
		return EmojiBank
	case "hotel":
		return EmojiHotel
	case "convenience_store":
		return EmojiConvenienceStore
	case "school":
		return EmojiSchool
	case "love_hotel":
		return EmojiLoveHotel
	case "wedding":
		return EmojiWedding
	case "classical_building":
		return EmojiClassicalBuilding
	case "church":
		return EmojiChurch
	case "mosque":
		return EmojiMosque
	case "synagogue":
		return EmojiSynagogue
	case "kaaba":
		return EmojiKaaba
	case "shinto_shrine":
		return EmojiShintoShrine
	case "japan":
		return EmojiJapan
	case "rice_scene":
		return EmojiRiceScene
	case "national_park":
		return EmojiNationalPark
	case "sunrise":
		return EmojiSunrise
	case "sunrise_over_mountains":
		return EmojiSunriseOverMountains
	case "stars":
		return EmojiStars
	case "sparkler":
		return EmojiSparkler
	case "fireworks":
		return EmojiFireworks
	case "city_sunrise":
		return EmojiCitySunrise
	case "city_sunset":
		return EmojiCitySunset
	case "cityscape":
		return EmojiCityscape
	case "night_with_stars":
		return EmojiNightWithStars
	case "milky_way":
		return EmojiMilkyWay
	case "bridge_at_night":
		return EmojiBridgeAtNight
	case "foggy":
		return EmojiFoggy
	case "watch":
		return EmojiWatch
	case "iphone":
		return EmojiIphone
	case "calling":
		return EmojiCalling
	case "computer":
		return EmojiComputer
	case "keyboard":
		return EmojiKeyboard
	case "desktop_computer":
		return EmojiDesktopComputer
	case "printer":
		return EmojiPrinter
	case "computer_mouse":
		return EmojiComputerMouse
	case "trackball":
		return EmojiTrackball
	case "joystick":
		return EmojiJoystick
	case "clamp":
		return EmojiClamp
	case "minidisc":
		return EmojiMinidisc
	case "floppy_disk":
		return EmojiFloppyDisk
	case "cd":
		return EmojiCd
	case "dvd":
		return EmojiDvd
	case "vhs":
		return EmojiVhs
	case "camera":
		return EmojiCamera
	case "camera_flash":
		return EmojiCameraFlash
	case "video_camera":
		return EmojiVideoCamera
	case "movie_camera":
		return EmojiMovieCamera
	case "film_projector":
		return EmojiFilmProjector
	case "film_strip":
		return EmojiFilmStrip
	case "telephone_receiver":
		return EmojiTelephoneReceiver
	case "phone", "telephone":
		return EmojiPhone
	case "pager":
		return EmojiPager
	case "fax":
		return EmojiFax
	case "tv":
		return EmojiTv
	case "radio":
		return EmojiRadio
	case "studio_microphone":
		return EmojiStudioMicrophone
	case "level_slider":
		return EmojiLevelSlider
	case "control_knobs":
		return EmojiControlKnobs
	case "stopwatch":
		return EmojiStopwatch
	case "timer_clock":
		return EmojiTimerClock
	case "alarm_clock":
		return EmojiAlarmClock
	case "mantelpiece_clock":
		return EmojiMantelpieceClock
	case "hourglass":
		return EmojiHourglass
	case "hourglass_flowing_sand":
		return EmojiHourglassFlowingSand
	case "satellite":
		return EmojiSatellite
	case "battery":
		return EmojiBattery
	case "electric_plug":
		return EmojiElectricPlug
	case "bulb":
		return EmojiBulb
	case "flashlight":
		return EmojiFlashlight
	case "candle":
		return EmojiCandle
	case "wastebasket":
		return EmojiWastebasket
	case "oil_drum":
		return EmojiOilDrum
	case "money_with_wings":
		return EmojiMoneyWithWings
	case "dollar":
		return EmojiDollar
	case "yen":
		return EmojiYen
	case "euro":
		return EmojiEuro
	case "pound":
		return EmojiPound
	case "moneybag":
		return EmojiMoneybag
	case "credit_card":
		return EmojiCreditCard
	case "gem":
		return EmojiGem
	case "balance_scale":
		return EmojiBalanceScale
	case "wrench":
		return EmojiWrench
	case "hammer":
		return EmojiHammer
	case "hammer_and_pick":
		return EmojiHammerAndPick
	case "hammer_and_wrench":
		return EmojiHammerAndWrench
	case "pick":
		return EmojiPick
	case "nut_and_bolt":
		return EmojiNutAndBolt
	case "gear":
		return EmojiGear
	case "chains":
		return EmojiChains
	case "gun":
		return EmojiGun
	case "bomb":
		return EmojiBomb
	case "hocho", "knife":
		return EmojiHocho
	case "dagger":
		return EmojiDagger
	case "crossed_swords":
		return EmojiCrossedSwords
	case "shield":
		return EmojiShield
	case "smoking":
		return EmojiSmoking
	case "coffin":
		return EmojiCoffin
	case "funeral_urn":
		return EmojiFuneralUrn
	case "amphora":
		return EmojiAmphora
	case "crystal_ball":
		return EmojiCrystalBall
	case "prayer_beads":
		return EmojiPrayerBeads
	case "barber":
		return EmojiBarber
	case "alembic":
		return EmojiAlembic
	case "telescope":
		return EmojiTelescope
	case "microscope":
		return EmojiMicroscope
	case "hole":
		return EmojiHole
	case "pill":
		return EmojiPill
	case "syringe":
		return EmojiSyringe
	case "thermometer":
		return EmojiThermometer
	case "toilet":
		return EmojiToilet
	case "potable_water":
		return EmojiPotableWater
	case "shower":
		return EmojiShower
	case "bathtub":
		return EmojiBathtub
	case "bath":
		return EmojiBath
	case "bellhop_bell":
		return EmojiBellhopBell
	case "key":
		return EmojiKey
	case "old_key":
		return EmojiOldKey
	case "door":
		return EmojiDoor
	case "couch_and_lamp":
		return EmojiCouchAndLamp
	case "bed":
		return EmojiBed
	case "sleeping_bed":
		return EmojiSleepingBed
	case "framed_picture":
		return EmojiFramedPicture
	case "shopping":
		return EmojiShopping
	case "shopping_cart":
		return EmojiShoppingCart
	case "gift":
		return EmojiGift
	case "balloon":
		return EmojiBalloon
	case "flags":
		return EmojiFlags
	case "ribbon":
		return EmojiRibbon
	case "confetti_ball":
		return EmojiConfettiBall
	case "tada":
		return EmojiTada
	case "dolls":
		return EmojiDolls
	case "izakaya_lantern", "lantern":
		return EmojiIzakayaLantern
	case "wind_chime":
		return EmojiWindChime
	case "email", "envelope":
		return EmojiEmail
	case "envelope_with_arrow":
		return EmojiEnvelopeWithArrow
	case "incoming_envelope":
		return EmojiIncomingEnvelope
	case "love_letter":
		return EmojiLoveLetter
	case "inbox_tray":
		return EmojiInboxTray
	case "outbox_tray":
		return EmojiOutboxTray
	case "package":
		return EmojiPackage
	case "label":
		return EmojiLabel
	case "mailbox_closed":
		return EmojiMailboxClosed
	case "mailbox":
		return EmojiMailbox
	case "mailbox_with_mail":
		return EmojiMailboxWithMail
	case "mailbox_with_no_mail":
		return EmojiMailboxWithNoMail
	case "postbox":
		return EmojiPostbox
	case "postal_horn":
		return EmojiPostalHorn
	case "scroll":
		return EmojiScroll
	case "page_with_curl":
		return EmojiPageWithCurl
	case "page_facing_up":
		return EmojiPageFacingUp
	case "bookmark_tabs":
		return EmojiBookmarkTabs
	case "bar_chart":
		return EmojiBarChart
	case "chart_with_upwards_trend":
		return EmojiChartWithUpwardsTrend
	case "chart_with_downwards_trend":
		return EmojiChartWithDownwardsTrend
	case "spiral_notepad":
		return EmojiSpiralNotepad
	case "spiral_calendar":
		return EmojiSpiralCalendar
	case "calendar":
		return EmojiCalendar
	case "date":
		return EmojiDate
	case "card_index":
		return EmojiCardIndex
	case "card_file_box":
		return EmojiCardFileBox
	case "ballot_box":
		return EmojiBallotBox
	case "file_cabinet":
		return EmojiFileCabinet
	case "clipboard":
		return EmojiClipboard
	case "file_folder":
		return EmojiFileFolder
	case "open_file_folder":
		return EmojiOpenFileFolder
	case "card_index_dividers":
		return EmojiCardIndexDividers
	case "newspaper_roll":
		return EmojiNewspaperRoll
	case "newspaper":
		return EmojiNewspaper
	case "notebook":
		return EmojiNotebook
	case "notebook_with_decorative_cover":
		return EmojiNotebookWithDecorativeCover
	case "ledger":
		return EmojiLedger
	case "closed_book":
		return EmojiClosedBook
	case "green_book":
		return EmojiGreenBook
	case "blue_book":
		return EmojiBlueBook
	case "orange_book":
		return EmojiOrangeBook
	case "books":
		return EmojiBooks
	case "book", "open_book":
		return EmojiBook
	case "bookmark":
		return EmojiBookmark
	case "link":
		return EmojiLink
	case "paperclip":
		return EmojiPaperclip
	case "paperclips":
		return EmojiPaperclips
	case "triangular_ruler":
		return EmojiTriangularRuler
	case "straight_ruler":
		return EmojiStraightRuler
	case "pushpin":
		return EmojiPushpin
	case "round_pushpin":
		return EmojiRoundPushpin
	case "scissors":
		return EmojiScissors
	case "pen":
		return EmojiPen
	case "fountain_pen":
		return EmojiFountainPen
	case "black_nib":
		return EmojiBlackNib
	case "paintbrush":
		return EmojiPaintbrush
	case "crayon":
		return EmojiCrayon
	case "memo", "pencil":
		return EmojiMemo
	case "pencil2":
		return EmojiPencil2
	case "mag":
		return EmojiMag
	case "mag_right":
		return EmojiMagRight
	case "lock_with_ink_pen":
		return EmojiLockWithInkPen
	case "closed_lock_with_key":
		return EmojiClosedLockWithKey
	case "lock":
		return EmojiLock
	case "unlock":
		return EmojiUnlock
	case "heart":
		return EmojiHeart
	case "yellow_heart":
		return EmojiYellowHeart
	case "green_heart":
		return EmojiGreenHeart
	case "blue_heart":
		return EmojiBlueHeart
	case "purple_heart":
		return EmojiPurpleHeart
	case "black_heart":
		return EmojiBlackHeart
	case "broken_heart":
		return EmojiBrokenHeart
	case "heavy_heart_exclamation":
		return EmojiHeavyHeartExclamation
	case "two_hearts":
		return EmojiTwoHearts
	case "revolving_hearts":
		return EmojiRevolvingHearts
	case "heartbeat":
		return EmojiHeartbeat
	case "heartpulse":
		return EmojiHeartpulse
	case "sparkling_heart":
		return EmojiSparklingHeart
	case "cupid":
		return EmojiCupid
	case "gift_heart":
		return EmojiGiftHeart
	case "heart_decoration":
		return EmojiHeartDecoration
	case "peace_symbol":
		return EmojiPeaceSymbol
	case "latin_cross":
		return EmojiLatinCross
	case "star_and_crescent":
		return EmojiStarAndCrescent
	case "om":
		return EmojiOm
	case "wheel_of_dharma":
		return EmojiWheelOfDharma
	case "star_of_david":
		return EmojiStarOfDavid
	case "six_pointed_star":
		return EmojiSixPointedStar
	case "menorah":
		return EmojiMenorah
	case "yin_yang":
		return EmojiYinYang
	case "orthodox_cross":
		return EmojiOrthodoxCross
	case "place_of_worship":
		return EmojiPlaceOfWorship
	case "ophiuchus":
		return EmojiOphiuchus
	case "aries":
		return EmojiAries
	case "taurus":
		return EmojiTaurus
	case "gemini":
		return EmojiGemini
	case "cancer":
		return EmojiCancer
	case "leo":
		return EmojiLeo
	case "virgo":
		return EmojiVirgo
	case "libra":
		return EmojiLibra
	case "scorpius":
		return EmojiScorpius
	case "sagittarius":
		return EmojiSagittarius
	case "capricorn":
		return EmojiCapricorn
	case "aquarius":
		return EmojiAquarius
	case "pisces":
		return EmojiPisces
	case "id":
		return EmojiId
	case "atom_symbol":
		return EmojiAtomSymbol
	case "accept":
		return EmojiAccept
	case "radioactive":
		return EmojiRadioactive
	case "biohazard":
		return EmojiBiohazard
	case "mobile_phone_off":
		return EmojiMobilePhoneOff
	case "vibration_mode":
		return EmojiVibrationMode
	case "u6709":
		return EmojiU6709
	case "u7121":
		return EmojiU7121
	case "u7533":
		return EmojiU7533
	case "u55b6":
		return EmojiU55b6
	case "u6708":
		return EmojiU6708
	case "eight_pointed_black_star":
		return EmojiEightPointedBlackStar
	case "vs":
		return EmojiVs
	case "white_flower":
		return EmojiWhiteFlower
	case "ideograph_advantage":
		return EmojiIdeographAdvantage
	case "secret":
		return EmojiSecret
	case "congratulations":
		return EmojiCongratulations
	case "u5408":
		return EmojiU5408
	case "u6e80":
		return EmojiU6e80
	case "u5272":
		return EmojiU5272
	case "u7981":
		return EmojiU7981
	case "a":
		return EmojiA
	case "b":
		return EmojiB
	case "ab":
		return EmojiAb
	case "cl":
		return EmojiCl
	case "o2":
		return EmojiO2
	case "sos":
		return EmojiSos
	case "x":
		return EmojiX
	case "o":
		return EmojiO
	case "stop_sign":
		return EmojiStopSign
	case "no_entry":
		return EmojiNoEntry
	case "name_badge":
		return EmojiNameBadge
	case "no_entry_sign":
		return EmojiNoEntrySign
	case "100":
		return Emoji100
	case "anger":
		return EmojiAnger
	case "hotsprings":
		return EmojiHotsprings
	case "no_pedestrians":
		return EmojiNoPedestrians
	case "do_not_litter":
		return EmojiDoNotLitter
	case "no_bicycles":
		return EmojiNoBicycles
	case "underage":
		return EmojiUnderage
	case "no_mobile_phones":
		return EmojiNoMobilePhones
	case "no_smoking":
		return EmojiNoSmoking
	case "exclamation", "heavy_exclamation_mark":
		return EmojiExclamation
	case "grey_exclamation":
		return EmojiGreyExclamation
	case "question":
		return EmojiQuestion
	case "grey_question":
		return EmojiGreyQuestion
	case "bangbang":
		return EmojiBangbang
	case "interrobang":
		return EmojiInterrobang
	case "low_brightness":
		return EmojiLowBrightness
	case "high_brightness":
		return EmojiHighBrightness
	case "part_alternation_mark":
		return EmojiPartAlternationMark
	case "warning":
		return EmojiWarning
	case "children_crossing":
		return EmojiChildrenCrossing
	case "trident":
		return EmojiTrident
	case "fleur_de_lis":
		return EmojiFleurDeLis
	case "beginner":
		return EmojiBeginner
	case "recycle":
		return EmojiRecycle
	case "white_check_mark":
		return EmojiWhiteCheckMark
	case "u6307":
		return EmojiU6307
	case "chart":
		return EmojiChart
	case "sparkle":
		return EmojiSparkle
	case "eight_spoked_asterisk":
		return EmojiEightSpokedAsterisk
	case "negative_squared_cross_mark":
		return EmojiNegativeSquaredCrossMark
	case "globe_with_meridians":
		return EmojiGlobeWithMeridians
	case "diamond_shape_with_a_dot_inside":
		return EmojiDiamondShapeWithADotInside
	case "m":
		return EmojiM
	case "cyclone":
		return EmojiCyclone
	case "zzz":
		return EmojiZzz
	case "atm":
		return EmojiAtm
	case "wc":
		return EmojiWc
	case "wheelchair":
		return EmojiWheelchair
	case "parking":
		return EmojiParking
	case "u7a7a":
		return EmojiU7a7a
	case "sa":
		return EmojiSa
	case "passport_control":
		return EmojiPassportControl
	case "customs":
		return EmojiCustoms
	case "baggage_claim":
		return EmojiBaggageClaim
	case "left_luggage":
		return EmojiLeftLuggage
	case "mens":
		return EmojiMens
	case "womens":
		return EmojiWomens
	case "baby_symbol":
		return EmojiBabySymbol
	case "restroom":
		return EmojiRestroom
	case "put_litter_in_its_place":
		return EmojiPutLitterInItsPlace
	case "cinema":
		return EmojiCinema
	case "signal_strength":
		return EmojiSignalStrength
	case "koko":
		return EmojiKoko
	case "symbols":
		return EmojiSymbols
	case "information_source":
		return EmojiInformationSource
	case "abc":
		return EmojiAbc
	case "abcd":
		return EmojiAbcd
	case "capital_abcd":
		return EmojiCapitalAbcd
	case "ng":
		return EmojiNg
	case "ok":
		return EmojiOk
	case "up":
		return EmojiUp
	case "cool":
		return EmojiCool
	case "new":
		return EmojiNew
	case "free":
		return EmojiFree
	case "zero":
		return EmojiZero
	case "one":
		return EmojiOne
	case "two":
		return EmojiTwo
	case "three":
		return EmojiThree
	case "four":
		return EmojiFour
	case "five":
		return EmojiFive
	case "six":
		return EmojiSix
	case "seven":
		return EmojiSeven
	case "eight":
		return EmojiEight
	case "nine":
		return EmojiNine
	case "keycap_ten":
		return EmojiKeycapTen
	case "1234":
		return Emoji1234
	case "hash":
		return EmojiHash
	case "asterisk":
		return EmojiAsterisk
	case "arrow_forward":
		return EmojiArrowForward
	case "pause_button":
		return EmojiPauseButton
	case "play_or_pause_button":
		return EmojiPlayOrPauseButton
	case "stop_button":
		return EmojiStopButton
	case "record_button":
		return EmojiRecordButton
	case "next_track_button":
		return EmojiNextTrackButton
	case "previous_track_button":
		return EmojiPreviousTrackButton
	case "fast_forward":
		return EmojiFastForward
	case "rewind":
		return EmojiRewind
	case "arrow_double_up":
		return EmojiArrowDoubleUp
	case "arrow_double_down":
		return EmojiArrowDoubleDown
	case "arrow_backward":
		return EmojiArrowBackward
	case "arrow_up_small":
		return EmojiArrowUpSmall
	case "arrow_down_small":
		return EmojiArrowDownSmall
	case "arrow_right":
		return EmojiArrowRight
	case "arrow_left":
		return EmojiArrowLeft
	case "arrow_up":
		return EmojiArrowUp
	case "arrow_down":
		return EmojiArrowDown
	case "arrow_upper_right":
		return EmojiArrowUpperRight
	case "arrow_lower_right":
		return EmojiArrowLowerRight
	case "arrow_lower_left":
		return EmojiArrowLowerLeft
	case "arrow_upper_left":
		return EmojiArrowUpperLeft
	case "arrow_up_down":
		return EmojiArrowUpDown
	case "left_right_arrow":
		return EmojiLeftRightArrow
	case "arrow_right_hook":
		return EmojiArrowRightHook
	case "leftwards_arrow_with_hook":
		return EmojiLeftwardsArrowWithHook
	case "arrow_heading_up":
		return EmojiArrowHeadingUp
	case "arrow_heading_down":
		return EmojiArrowHeadingDown
	case "twisted_rightwards_arrows":
		return EmojiTwistedRightwardsArrows
	case "repeat":
		return EmojiRepeat
	case "repeat_one":
		return EmojiRepeatOne
	case "arrows_counterclockwise":
		return EmojiArrowsCounterclockwise
	case "arrows_clockwise":
		return EmojiArrowsClockwise
	case "musical_note":
		return EmojiMusicalNote
	case "notes":
		return EmojiNotes
	case "heavy_plus_sign":
		return EmojiHeavyPlusSign
	case "heavy_minus_sign":
		return EmojiHeavyMinusSign
	case "heavy_division_sign":
		return EmojiHeavyDivisionSign
	case "heavy_multiplication_x":
		return EmojiHeavyMultiplicationX
	case "heavy_dollar_sign":
		return EmojiHeavyDollarSign
	case "currency_exchange":
		return EmojiCurrencyExchange
	case "tm":
		return EmojiTm
	case "copyright":
		return EmojiCopyright
	case "registered":
		return EmojiRegistered
	case "wavy_dash":
		return EmojiWavyDash
	case "curly_loop":
		return EmojiCurlyLoop
	case "loop":
		return EmojiLoop
	case "end":
		return EmojiEnd
	case "back":
		return EmojiBack
	case "on":
		return EmojiOn
	case "top":
		return EmojiTop
	case "soon":
		return EmojiSoon
	case "heavy_check_mark":
		return EmojiHeavyCheckMark
	case "ballot_box_with_check":
		return EmojiBallotBoxWithCheck
	case "radio_button":
		return EmojiRadioButton
	case "white_circle":
		return EmojiWhiteCircle
	case "black_circle":
		return EmojiBlackCircle
	case "red_circle":
		return EmojiRedCircle
	case "large_blue_circle":
		return EmojiLargeBlueCircle
	case "small_red_triangle":
		return EmojiSmallRedTriangle
	case "small_red_triangle_down":
		return EmojiSmallRedTriangleDown
	case "small_orange_diamond":
		return EmojiSmallOrangeDiamond
	case "small_blue_diamond":
		return EmojiSmallBlueDiamond
	case "large_orange_diamond":
		return EmojiLargeOrangeDiamond
	case "large_blue_diamond":
		return EmojiLargeBlueDiamond
	case "white_square_button":
		return EmojiWhiteSquareButton
	case "black_square_button":
		return EmojiBlackSquareButton
	case "black_small_square":
		return EmojiBlackSmallSquare
	case "white_small_square":
		return EmojiWhiteSmallSquare
	case "black_medium_small_square":
		return EmojiBlackMediumSmallSquare
	case "white_medium_small_square":
		return EmojiWhiteMediumSmallSquare
	case "black_medium_square":
		return EmojiBlackMediumSquare
	case "white_medium_square":
		return EmojiWhiteMediumSquare
	case "black_large_square":
		return EmojiBlackLargeSquare
	case "white_large_square":
		return EmojiWhiteLargeSquare
	case "speaker":
		return EmojiSpeaker
	case "mute":
		return EmojiMute
	case "sound":
		return EmojiSound
	case "loud_sound":
		return EmojiLoudSound
	case "bell":
		return EmojiBell
	case "no_bell":
		return EmojiNoBell
	case "mega":
		return EmojiMega
	case "loudspeaker":
		return EmojiLoudspeaker
	case "eye_speech_bubble":
		return EmojiEyeSpeechBubble
	case "speech_balloon":
		return EmojiSpeechBalloon
	case "thought_balloon":
		return EmojiThoughtBalloon
	case "right_anger_bubble":
		return EmojiRightAngerBubble
	case "spades":
		return EmojiSpades
	case "clubs":
		return EmojiClubs
	case "hearts":
		return EmojiHearts
	case "diamonds":
		return EmojiDiamonds
	case "black_joker":
		return EmojiBlackJoker
	case "flower_playing_cards":
		return EmojiFlowerPlayingCards
	case "mahjong":
		return EmojiMahjong
	case "clock1":
		return EmojiClock1
	case "clock2":
		return EmojiClock2
	case "clock3":
		return EmojiClock3
	case "clock4":
		return EmojiClock4
	case "clock5":
		return EmojiClock5
	case "clock6":
		return EmojiClock6
	case "clock7":
		return EmojiClock7
	case "clock8":
		return EmojiClock8
	case "clock9":
		return EmojiClock9
	case "clock10":
		return EmojiClock10
	case "clock11":
		return EmojiClock11
	case "clock12":
		return EmojiClock12
	case "clock130":
		return EmojiClock130
	case "clock230":
		return EmojiClock230
	case "clock330":
		return EmojiClock330
	case "clock430":
		return EmojiClock430
	case "clock530":
		return EmojiClock530
	case "clock630":
		return EmojiClock630
	case "clock730":
		return EmojiClock730
	case "clock830":
		return EmojiClock830
	case "clock930":
		return EmojiClock930
	case "clock1030":
		return EmojiClock1030
	case "clock1130":
		return EmojiClock1130
	case "clock1230":
		return EmojiClock1230
	case "white_flag":
		return EmojiWhiteFlag
	case "black_flag":
		return EmojiBlackFlag
	case "checkered_flag":
		return EmojiCheckeredFlag
	case "triangular_flag_on_post":
		return EmojiTriangularFlagOnPost
	case "rainbow_flag":
		return EmojiRainbowFlag
	case "afghanistan":
		return EmojiAfghanistan
	case "aland_islands":
		return EmojiAlandIslands
	case "albania":
		return EmojiAlbania
	case "algeria":
		return EmojiAlgeria
	case "american_samoa":
		return EmojiAmericanSamoa
	case "andorra":
		return EmojiAndorra
	case "angola":
		return EmojiAngola
	case "anguilla":
		return EmojiAnguilla
	case "antarctica":
		return EmojiAntarctica
	case "antigua_barbuda":
		return EmojiAntiguaBarbuda
	case "argentina":
		return EmojiArgentina
	case "armenia":
		return EmojiArmenia
	case "aruba":
		return EmojiAruba
	case "australia":
		return EmojiAustralia
	case "austria":
		return EmojiAustria
	case "azerbaijan":
		return EmojiAzerbaijan
	case "bahamas":
		return EmojiBahamas
	case "bahrain":
		return EmojiBahrain
	case "bangladesh":
		return EmojiBangladesh
	case "barbados":
		return EmojiBarbados
	case "belarus":
		return EmojiBelarus
	case "belgium":
		return EmojiBelgium
	case "belize":
		return EmojiBelize
	case "benin":
		return EmojiBenin
	case "bermuda":
		return EmojiBermuda
	case "bhutan":
		return EmojiBhutan
	case "bolivia":
		return EmojiBolivia
	case "caribbean_netherlands":
		return EmojiCaribbeanNetherlands
	case "bosnia_herzegovina":
		return EmojiBosniaHerzegovina
	case "botswana":
		return EmojiBotswana
	case "brazil":
		return EmojiBrazil
	case "british_indian_ocean_territory":
		return EmojiBritishIndianOceanTerritory
	case "british_virgin_islands":
		return EmojiBritishVirginIslands
	case "brunei":
		return EmojiBrunei
	case "bulgaria":
		return EmojiBulgaria
	case "burkina_faso":
		return EmojiBurkinaFaso
	case "burundi":
		return EmojiBurundi
	case "cape_verde":
		return EmojiCapeVerde
	case "cambodia":
		return EmojiCambodia
	case "cameroon":
		return EmojiCameroon
	case "canada":
		return EmojiCanada
	case "canary_islands":
		return EmojiCanaryIslands
	case "cayman_islands":
		return EmojiCaymanIslands
	case "central_african_republic":
		return EmojiCentralAfricanRepublic
	case "chad":
		return EmojiChad
	case "chile":
		return EmojiChile
	case "cn":
		return EmojiCn
	case "christmas_island":
		return EmojiChristmasIsland
	case "cocos_islands":
		return EmojiCocosIslands
	case "colombia":
		return EmojiColombia
	case "comoros":
		return EmojiComoros
	case "congo_brazzaville":
		return EmojiCongoBrazzaville
	case "congo_kinshasa":
		return EmojiCongoKinshasa
	case "cook_islands":
		return EmojiCookIslands
	case "costa_rica":
		return EmojiCostaRica
	case "cote_divoire":
		return EmojiCoteDivoire
	case "croatia":
		return EmojiCroatia
	case "cuba":
		return EmojiCuba
	case "curacao":
		return EmojiCuracao
	case "cyprus":
		return EmojiCyprus
	case "czech_republic":
		return EmojiCzechRepublic
	case "denmark":
		return EmojiDenmark
	case "djibouti":
		return EmojiDjibouti
	case "dominica":
		return EmojiDominica
	case "dominican_republic":
		return EmojiDominicanRepublic
	case "ecuador":
		return EmojiEcuador
	case "egypt":
		return EmojiEgypt
	case "el_salvador":
		return EmojiElSalvador
	case "equatorial_guinea":
		return EmojiEquatorialGuinea
	case "eritrea":
		return EmojiEritrea
	case "estonia":
		return EmojiEstonia
	case "ethiopia":
		return EmojiEthiopia
	case "eu", "european_union":
		return EmojiEu
	case "falkland_islands":
		return EmojiFalklandIslands
	case "faroe_islands":
		return EmojiFaroeIslands
	case "fiji":
		return EmojiFiji
	case "finland":
		return EmojiFinland
	case "fr":
		return EmojiFr
	case "french_guiana":
		return EmojiFrenchGuiana
	case "french_polynesia":
		return EmojiFrenchPolynesia
	case "french_southern_territories":
		return EmojiFrenchSouthernTerritories
	case "gabon":
		return EmojiGabon
	case "gambia":
		return EmojiGambia
	case "georgia":
		return EmojiGeorgia
	case "de":
		return EmojiDe
	case "ghana":
		return EmojiGhana
	case "gibraltar":
		return EmojiGibraltar
	case "greece":
		return EmojiGreece
	case "greenland":
		return EmojiGreenland
	case "grenada":
		return EmojiGrenada
	case "guadeloupe":
		return EmojiGuadeloupe
	case "guam":
		return EmojiGuam
	case "guatemala":
		return EmojiGuatemala
	case "guernsey":
		return EmojiGuernsey
	case "guinea":
		return EmojiGuinea
	case "guinea_bissau":
		return EmojiGuineaBissau
	case "guyana":
		return EmojiGuyana
	case "haiti":
		return EmojiHaiti
	case "honduras":
		return EmojiHonduras
	case "hong_kong":
		return EmojiHongKong
	case "hungary":
		return EmojiHungary
	case "iceland":
		return EmojiIceland
	case "india":
		return EmojiIndia
	case "indonesia":
		return EmojiIndonesia
	case "iran":
		return EmojiIran
	case "iraq":
		return EmojiIraq
	case "ireland":
		return EmojiIreland
	case "isle_of_man":
		return EmojiIsleOfMan
	case "israel":
		return EmojiIsrael
	case "it":
		return EmojiIt
	case "jamaica":
		return EmojiJamaica
	case "jp":
		return EmojiJp
	case "crossed_flags":
		return EmojiCrossedFlags
	case "jersey":
		return EmojiJersey
	case "jordan":
		return EmojiJordan
	case "kazakhstan":
		return EmojiKazakhstan
	case "kenya":
		return EmojiKenya
	case "kiribati":
		return EmojiKiribati
	case "kosovo":
		return EmojiKosovo
	case "kuwait":
		return EmojiKuwait
	case "kyrgyzstan":
		return EmojiKyrgyzstan
	case "laos":
		return EmojiLaos
	case "latvia":
		return EmojiLatvia
	case "lebanon":
		return EmojiLebanon
	case "lesotho":
		return EmojiLesotho
	case "liberia":
		return EmojiLiberia
	case "libya":
		return EmojiLibya
	case "liechtenstein":
		return EmojiLiechtenstein
	case "lithuania":
		return EmojiLithuania
	case "luxembourg":
		return EmojiLuxembourg
	case "macau":
		return EmojiMacau
	case "macedonia":
		return EmojiMacedonia
	case "madagascar":
		return EmojiMadagascar
	case "malawi":
		return EmojiMalawi
	case "malaysia":
		return EmojiMalaysia
	case "maldives":
		return EmojiMaldives
	case "mali":
		return EmojiMali
	case "malta":
		return EmojiMalta
	case "marshall_islands":
		return EmojiMarshallIslands
	case "martinique":
		return EmojiMartinique
	case "mauritania":
		return EmojiMauritania
	case "mauritius":
		return EmojiMauritius
	case "mayotte":
		return EmojiMayotte
	case "mexico":
		return EmojiMexico
	case "micronesia":
		return EmojiMicronesia
	case "moldova":
		return EmojiMoldova
	case "monaco":
		return EmojiMonaco
	case "mongolia":
		return EmojiMongolia
	case "montenegro":
		return EmojiMontenegro
	case "montserrat":
		return EmojiMontserrat
	case "morocco":
		return EmojiMorocco
	case "mozambique":
		return EmojiMozambique
	case "myanmar":
		return EmojiMyanmar
	case "namibia":
		return EmojiNamibia
	case "nauru":
		return EmojiNauru
	case "nepal":
		return EmojiNepal
	case "netherlands":
		return EmojiNetherlands
	case "new_caledonia":
		return EmojiNewCaledonia
	case "new_zealand":
		return EmojiNewZealand
	case "nicaragua":
		return EmojiNicaragua
	case "niger":
		return EmojiNiger
	case "nigeria":
		return EmojiNigeria
	case "niue":
		return EmojiNiue
	case "norfolk_island":
		return EmojiNorfolkIsland
	case "northern_mariana_islands":
		return EmojiNorthernMarianaIslands
	case "north_korea":
		return EmojiNorthKorea
	case "norway":
		return EmojiNorway
	case "oman":
		return EmojiOman
	case "pakistan":
		return EmojiPakistan
	case "palau":
		return EmojiPalau
	case "palestinian_territories":
		return EmojiPalestinianTerritories
	case "panama":
		return EmojiPanama
	case "papua_new_guinea":
		return EmojiPapuaNewGuinea
	case "paraguay":
		return EmojiParaguay
	case "peru":
		return EmojiPeru
	case "philippines":
		return EmojiPhilippines
	case "pitcairn_islands":
		return EmojiPitcairnIslands
	case "poland":
		return EmojiPoland
	case "portugal":
		return EmojiPortugal
	case "puerto_rico":
		return EmojiPuertoRico
	case "qatar":
		return EmojiQatar
	case "reunion":
		return EmojiReunion
	case "romania":
		return EmojiRomania
	case "ru":
		return EmojiRu
	case "rwanda":
		return EmojiRwanda
	case "st_barthelemy":
		return EmojiStBarthelemy
	case "st_helena":
		return EmojiStHelena
	case "st_kitts_nevis":
		return EmojiStKittsNevis
	case "st_lucia":
		return EmojiStLucia
	case "st_pierre_miquelon":
		return EmojiStPierreMiquelon
	case "st_vincent_grenadines":
		return EmojiStVincentGrenadines
	case "samoa":
		return EmojiSamoa
	case "san_marino":
		return EmojiSanMarino
	case "sao_tome_principe":
		return EmojiSaoTomePrincipe
	case "saudi_arabia":
		return EmojiSaudiArabia
	case "senegal":
		return EmojiSenegal
	case "serbia":
		return EmojiSerbia
	case "seychelles":
		return EmojiSeychelles
	case "sierra_leone":
		return EmojiSierraLeone
	case "singapore":
		return EmojiSingapore
	case "sint_maarten":
		return EmojiSintMaarten
	case "slovakia":
		return EmojiSlovakia
	case "slovenia":
		return EmojiSlovenia
	case "solomon_islands":
		return EmojiSolomonIslands
	case "somalia":
		return EmojiSomalia
	case "south_africa":
		return EmojiSouthAfrica
	case "south_georgia_south_sandwich_islands":
		return EmojiSouthGeorgiaSouthSandwichIslands
	case "kr":
		return EmojiKr
	case "south_sudan":
		return EmojiSouthSudan
	case "es":
		return EmojiEs
	case "sri_lanka":
		return EmojiSriLanka
	case "sudan":
		return EmojiSudan
	case "suriname":
		return EmojiSuriname
	case "swaziland":
		return EmojiSwaziland
	case "sweden":
		return EmojiSweden
	case "switzerland":
		return EmojiSwitzerland
	case "syria":
		return EmojiSyria
	case "taiwan":
		return EmojiTaiwan
	case "tajikistan":
		return EmojiTajikistan
	case "tanzania":
		return EmojiTanzania
	case "thailand":
		return EmojiThailand
	case "timor_leste":
		return EmojiTimorLeste
	case "togo":
		return EmojiTogo
	case "tokelau":
		return EmojiTokelau
	case "tonga":
		return EmojiTonga
	case "trinidad_tobago":
		return EmojiTrinidadTobago
	case "tunisia":
		return EmojiTunisia
	case "tr":
		return EmojiTr
	case "turkmenistan":
		return EmojiTurkmenistan
	case "turks_caicos_islands":
		return EmojiTurksCaicosIslands
	case "tuvalu":
		return EmojiTuvalu
	case "uganda":
		return EmojiUganda
	case "ukraine":
		return EmojiUkraine
	case "united_arab_emirates":
		return EmojiUnitedArabEmirates
	case "gb", "uk":
		return EmojiGb
	case "us":
		return EmojiUs
	case "us_virgin_islands":
		return EmojiUsVirginIslands
	case "uruguay":
		return EmojiUruguay
	case "uzbekistan":
		return EmojiUzbekistan
	case "vanuatu":
		return EmojiVanuatu
	case "vatican_city":
		return EmojiVaticanCity
	case "venezuela":
		return EmojiVenezuela
	case "vietnam":
		return EmojiVietnam
	case "wallis_futuna":
		return EmojiWallisFutuna
	case "western_sahara":
		return EmojiWesternSahara
	case "yemen":
		return EmojiYemen
	case "zambia":
		return EmojiZambia
	case "zimbabwe":
		return EmojiZimbabwe
	case "basecamp":
		return EmojiBasecamp
	case "basecampy":
		return EmojiBasecampy
	case "bowtie":
		return EmojiBowtie
	case "feelsgood":
		return EmojiFeelsgood
	case "finnadie":
		return EmojiFinnadie
	case "goberserk":
		return EmojiGoberserk
	case "godmode":
		return EmojiGodmode
	case "hurtrealbad":
		return EmojiHurtrealbad
	case "neckbeard":
		return EmojiNeckbeard
	case "octocat":
		return EmojiOctocat
	case "rage1":
		return EmojiRage1
	case "rage2":
		return EmojiRage2
	case "rage3":
		return EmojiRage3
	case "rage4":
		return EmojiRage4
	case "shipit", "squirrel":
		return EmojiShipit
	case "suspect":
		return EmojiSuspect
	case "trollface":
		return EmojiTrollface
	}
	return ""
}
