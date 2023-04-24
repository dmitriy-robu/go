package controllers

type Controllers struct {
	CaseBattleController CaseBattleController
	UserController       UserController
	ReferralController   ReferralController
	//ProvablyFairController ProvablyFairController
	BoxController       BoxController
	SteamAuthController SteamAuthController
}

func NewControllers(
	caseBattleController CaseBattleController,
	userController UserController,
	referralController ReferralController,
	//provablyFairController ProvablyFairController,
	boxController BoxController,
	steamAuthController SteamAuthController,
) Controllers {
	return Controllers{
		CaseBattleController: caseBattleController,
		UserController:       userController,
		ReferralController:   referralController,
		//ProvablyFairController: provablyFairController,
		BoxController:       boxController,
		SteamAuthController: steamAuthController,
	}
}
