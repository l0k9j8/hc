package main

import (
    "time"
    
    "github.com/brutella/log"
    "github.com/brutella/hap/app"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
)

func main() {
    log.Info = false
    
    conf := app.NewConfig()
    conf.DatabaseDir = "./data"
    
    app, err := app.NewApp(conf)
    if err != nil {
        log.Fatal(err)
    }
    
    info := model.Info{
        Name: "My Switch",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Switchy",
    }
    
    sw := accessory.NewSwitch(info)
    sw.OnStateChanged(func(on bool) {
        if on == true {
            log.Println("[INFO] Switch on")
        } else {
            log.Println("[INFO] Switch off")
        }
    })
    
    app.AddAccessory(sw.Accessory)
    
    go func() {
        timer := time.NewTimer(2 * time.Second)
        for {
            <- timer.C
            log.Println("[INFO] Update switch")
            sw.SetOn(sw.IsOn() == false)
            timer.Reset(2 * time.Second)
        }
    }()
    
    app.Run()
}