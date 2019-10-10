package config


//type Config struct {
//	Name string
//}
//
//func ConfInit(cfg string) error {
//	c := Config{
//		Name: cfg,
//	}
//
//	// 初始化配置文件
//	if err := c.initConfig(); err != nil {
//		return err
//	}
//
//	// 监控配置文件变化并热加载程序
//	//c.watchConfig()
//
//	return nil
//}
//
//func (c *Config) initConfig() error {
//	_ = viper.AddRemoteProvider("etcd", "http://127.0.0.1:2379", "/conf/hugo.yml")
//	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
//
//	// read from remote config the first time.
//	if err := viper.ReadRemoteConfig(); err != nil { // viper解析配置文件
//		return err
//	}
//
//	//// unmarshal config
//	//_ = viper.Unmarshal(&runtime_conf)
//
//	// open a goroutine to watch remote changes forever
//	go func(){
//		for {
//			time.Sleep(time.Second * 5) // delay after each request
//
//			// currently, only tested with etcd support
//			err := viper.WatchRemoteConfig()
//			if err != nil {
//				log.Errorf("unable to read remote config: %v", err)
//				continue
//			}
//
//			// unmarshal new config into our runtime config struct. you can also use channel
//			// to implement a signal to notify the system of the changes
//			//viper.Unmarshal(&runtime_conf)
//		}
//	}()
//
//
//	return nil
//}
//
//// 监控配置文件变化并热加载程序
////func (c *Config) watchConfig() {
////	viper.WatchConfig()
////	viper.OnConfigChange(func(e fsnotify.Event) {
////		logger.HandlerLogger().Info("Config file changed: ", e.Name)
////	})
////}
