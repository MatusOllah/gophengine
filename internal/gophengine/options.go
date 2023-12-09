package gophengine

var Options struct {
	PC4R          bool   `long:"pc4r" description:"Pééčivoo chleba 4 rohlíky"`
	ExtractAssets bool   `long:"extract-assets" description:"Extract embedded assets"`
	Config        string `long:"config" description:"Path to config.gecfg"`
	Progress      string `long:"progress" description:"Path to progress.gecfg"`
	VSync         bool   `long:"vsync" description:"Enable VSync"`
}
