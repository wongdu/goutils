package cron

func ClearOldOssFiles(){
for {
		time.Sleep(time.Hour * 5)
		cleanStale()
	}

}

func cleanStale(){


}