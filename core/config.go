package core

import (
  "fmt"
  "flag"
  "os"
  "encoding/json"
  "github.com/webability-go/xconfig"
  "github.com/webability-go/xamboo/utils"
)

type Listener struct {
  Name string `json:"name"`
  IP string `json:"ip"`
  Port string `json:"port"`
  Protocol string `json:"protocol"`
  ReadTimeOut int `json:"readtimeout"`
  WriteTimeOut int `json:"writetimeout"`
  HeaderSize int `json:"headersize"`
}

type Host struct {
  Name string `json:"name"`
  Listeners []string `json:"listeners"`
  HostNames []string `json:"hostnames"`
  Cert string `json:"cert"`
  PrivateKey string `json:"key"`
  ConfigFile []string `json:"config"`
  Config *xconfig.XConfig
}

type ConfigDef struct {
    File string
    Listeners []Listener
    Hosts []Host
}

var Config = &ConfigDef{}

// Then main xamboo runner
func (c *ConfigDef) Load() error {
  flag.StringVar(&c.File, "config", "", "configuration file")
  flag.Parse()
  
  if c.File == "" {
    return nil
  }
  configFile, err := os.Open(c.File)
  defer configFile.Close()
  if err != nil {
    return err
  }

  jsonParser := json.NewDecoder(configFile)
  if err = jsonParser.Decode(c); err != nil {
    return err
  }

  // Parse the XConfig file
  if c.Hosts != nil {
    for i, _ := range c.Hosts {
      if c.Hosts[i].ConfigFile != nil {
        for j, _ := range c.Hosts[i].ConfigFile {
          lc, _ := xconfig.Load(c.Hosts[i].ConfigFile[j])
          fmt.Printf("%p\n", lc)
          
          c.Hosts[i].Config = lc
          fmt.Println(c.Hosts[i].Config)
          fmt.Printf("%p\n", c.Hosts[i].Config)
        }
      }
    }
  }
  
  // Load the configuration file
  fmt.Println("Config loaded " + c.File)

  fmt.Println("CONFIG DENTRO DE CONFIG.GO:")
  fmt.Println(c.Hosts[0].Config)
  fmt.Printf("%p\n", c.Hosts[0].Config)

  fmt.Printf("%+v\n", c)
  return nil
}

func (c *ConfigDef) GetListener(host string, port string, secure bool) (*Host, *Listener) {

  for _, h := range c.Hosts {
    if utils.SearchInArray(host, h.HostNames) {
      // search the actual active listener
      
      
      
      
      return &h, &c.Listeners[0]
    }
  }
  return nil, nil
}
