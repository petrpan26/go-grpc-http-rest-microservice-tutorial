{"filter":false,"title":"main.go","tooltip":"/main.go","undoManager":{"mark":100,"position":100,"stack":[[{"start":{"row":14,"column":1},"end":{"row":14,"column":2},"action":"insert","lines":["u"],"id":34},{"start":{"row":14,"column":2},"end":{"row":14,"column":3},"action":"insert","lines":["u"]},{"start":{"row":14,"column":3},"end":{"row":14,"column":4},"action":"insert","lines":["i"]}],[{"start":{"row":14,"column":1},"end":{"row":14,"column":4},"action":"remove","lines":["uui"],"id":35},{"start":{"row":14,"column":1},"end":{"row":14,"column":5},"action":"insert","lines":["uuid"]}],[{"start":{"row":14,"column":1},"end":{"row":14,"column":5},"action":"remove","lines":["uuid"],"id":36}],[{"start":{"row":14,"column":1},"end":{"row":14,"column":2},"action":"insert","lines":["p"],"id":37},{"start":{"row":14,"column":2},"end":{"row":14,"column":3},"action":"insert","lines":["t"]}],[{"start":{"row":14,"column":1},"end":{"row":14,"column":3},"action":"remove","lines":["pt"],"id":38},{"start":{"row":14,"column":1},"end":{"row":14,"column":7},"action":"insert","lines":["ptypes"]}],[{"start":{"row":14,"column":7},"end":{"row":14,"column":8},"action":"insert","lines":["."],"id":39}],[{"start":{"row":14,"column":7},"end":{"row":14,"column":8},"action":"remove","lines":["."],"id":40}],[{"start":{"row":14,"column":1},"end":{"row":14,"column":7},"action":"remove","lines":["ptypes"],"id":41}],[{"start":{"row":14,"column":1},"end":{"row":14,"column":2},"action":"insert","lines":["u"],"id":42},{"start":{"row":14,"column":2},"end":{"row":14,"column":3},"action":"insert","lines":["u"]},{"start":{"row":14,"column":3},"end":{"row":14,"column":4},"action":"insert","lines":["i"]}],[{"start":{"row":14,"column":1},"end":{"row":14,"column":4},"action":"remove","lines":["uui"],"id":43},{"start":{"row":14,"column":1},"end":{"row":14,"column":5},"action":"insert","lines":["uuid"]}],[{"start":{"row":14,"column":5},"end":{"row":14,"column":6},"action":"insert","lines":["."],"id":44}],[{"start":{"row":9,"column":13},"end":{"row":14,"column":8},"action":"remove","lines":["","\tvar val = ptypes.TimestampNow()","\tval = ptypes.TimestampNow()","\tfmt.Print(val)","\tclient, err := as.NewClient()","\tptypes."],"id":45}],[{"start":{"row":4,"column":43},"end":{"row":6,"column":25},"action":"remove","lines":["","\t\"github.com/golang/protobuf/ptypes\"","\t\"github.com/google/uuid\""],"id":46}],[{"start":{"row":7,"column":13},"end":{"row":8,"column":0},"action":"insert","lines":["",""],"id":47},{"start":{"row":8,"column":0},"end":{"row":8,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":8,"column":1},"end":{"row":9,"column":0},"action":"insert","lines":["client, err := as.NewClient(\"192.168.1.150\", 3000)",""],"id":48}],[{"start":{"row":8,"column":51},"end":{"row":9,"column":0},"action":"remove","lines":["",""],"id":49}],[{"start":{"row":8,"column":30},"end":{"row":8,"column":43},"action":"remove","lines":["192.168.1.150"],"id":50},{"start":{"row":8,"column":30},"end":{"row":8,"column":40},"action":"insert","lines":["172.17.0.2"]}],[{"start":{"row":8,"column":48},"end":{"row":9,"column":0},"action":"insert","lines":["",""],"id":51},{"start":{"row":9,"column":0},"end":{"row":9,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":9,"column":1},"end":{"row":30,"column":14},"action":"insert","lines":["key, err := as.NewKey(\"namespace\", \"set\",","    \"key value goes here and can be any supported primitive\")","if err != nil {","    log.Fatal(err)","}","","bin1 := as.NewBin(\"bin1\", \"value1\")","bin2 := as.NewBin(\"bin2\", \"value2\")","","// Write a record","err = client.PutBins(nil, key, bin1, bin2)","if err != nil {","    log.Fatal(err)","}","","// Read a record","record, err := client.Get(nil, key)","if err != nil {","    log.Fatal(err)","}","","client.Close()"],"id":52}],[{"start":{"row":10,"column":0},"end":{"row":10,"column":4},"action":"remove","lines":["    "],"id":53,"ignore":true},{"start":{"row":10,"column":0},"end":{"row":10,"column":2},"action":"insert","lines":["\t\t"]},{"start":{"row":11,"column":0},"end":{"row":11,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":12,"column":0},"end":{"row":12,"column":4},"action":"remove","lines":["    "]},{"start":{"row":12,"column":0},"end":{"row":12,"column":2},"action":"insert","lines":["\t\t"]},{"start":{"row":13,"column":0},"end":{"row":13,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":15,"column":0},"end":{"row":15,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":16,"column":0},"end":{"row":16,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":18,"column":0},"end":{"row":18,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":19,"column":0},"end":{"row":19,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":20,"column":0},"end":{"row":20,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":21,"column":0},"end":{"row":21,"column":4},"action":"remove","lines":["    "]},{"start":{"row":21,"column":0},"end":{"row":21,"column":2},"action":"insert","lines":["\t\t"]},{"start":{"row":22,"column":0},"end":{"row":22,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":24,"column":0},"end":{"row":24,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":25,"column":0},"end":{"row":25,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":26,"column":0},"end":{"row":26,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":27,"column":0},"end":{"row":27,"column":4},"action":"remove","lines":["    "]},{"start":{"row":27,"column":0},"end":{"row":27,"column":2},"action":"insert","lines":["\t\t"]},{"start":{"row":28,"column":0},"end":{"row":28,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":30,"column":0},"end":{"row":30,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":2,"column":8},"end":{"row":3,"column":6},"action":"remove","lines":["","\t\"fmt\""],"id":54}],[{"start":{"row":3,"column":1},"end":{"row":3,"column":2},"action":"insert","lines":["a"],"id":55},{"start":{"row":3,"column":2},"end":{"row":3,"column":3},"action":"insert","lines":["s"]}],[{"start":{"row":3,"column":3},"end":{"row":3,"column":4},"action":"insert","lines":[" "],"id":56}],[{"start":{"row":3,"column":46},"end":{"row":4,"column":0},"action":"insert","lines":["",""],"id":57},{"start":{"row":4,"column":0},"end":{"row":4,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":4,"column":1},"end":{"row":4,"column":3},"action":"insert","lines":["\"\""],"id":58}],[{"start":{"row":4,"column":2},"end":{"row":4,"column":3},"action":"insert","lines":["l"],"id":59},{"start":{"row":4,"column":3},"end":{"row":4,"column":4},"action":"insert","lines":["o"]},{"start":{"row":4,"column":4},"end":{"row":4,"column":5},"action":"insert","lines":["g"]}],[{"start":{"row":28,"column":2},"end":{"row":29,"column":0},"action":"insert","lines":["",""],"id":60},{"start":{"row":29,"column":0},"end":{"row":29,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":4,"column":6},"end":{"row":5,"column":0},"action":"insert","lines":["",""],"id":61},{"start":{"row":5,"column":0},"end":{"row":5,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":5,"column":0},"end":{"row":5,"column":1},"action":"remove","lines":["\t"],"id":62},{"start":{"row":4,"column":6},"end":{"row":5,"column":0},"action":"remove","lines":["",""]}],[{"start":{"row":29,"column":1},"end":{"row":30,"column":0},"action":"insert","lines":["",""],"id":63},{"start":{"row":30,"column":0},"end":{"row":30,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":30,"column":0},"end":{"row":30,"column":1},"action":"remove","lines":["\t"],"id":64},{"start":{"row":29,"column":1},"end":{"row":30,"column":0},"action":"remove","lines":["",""]}],[{"start":{"row":29,"column":1},"end":{"row":30,"column":0},"action":"insert","lines":["",""],"id":65},{"start":{"row":30,"column":0},"end":{"row":30,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":30,"column":1},"end":{"row":30,"column":2},"action":"insert","lines":["l"]},{"start":{"row":30,"column":2},"end":{"row":30,"column":3},"action":"insert","lines":["o"]},{"start":{"row":30,"column":3},"end":{"row":30,"column":4},"action":"insert","lines":["g"]},{"start":{"row":30,"column":4},"end":{"row":30,"column":5},"action":"insert","lines":["."]}],[{"start":{"row":30,"column":5},"end":{"row":30,"column":6},"action":"insert","lines":["P"],"id":66}],[{"start":{"row":30,"column":5},"end":{"row":30,"column":6},"action":"remove","lines":["P"],"id":67}],[{"start":{"row":30,"column":5},"end":{"row":30,"column":6},"action":"insert","lines":["P"],"id":68}],[{"start":{"row":30,"column":5},"end":{"row":30,"column":6},"action":"remove","lines":["P"],"id":69},{"start":{"row":30,"column":5},"end":{"row":30,"column":12},"action":"insert","lines":["Print()"]}],[{"start":{"row":30,"column":11},"end":{"row":30,"column":12},"action":"insert","lines":["r"],"id":70},{"start":{"row":30,"column":12},"end":{"row":30,"column":13},"action":"insert","lines":["e"]}],[{"start":{"row":30,"column":11},"end":{"row":30,"column":13},"action":"remove","lines":["re"],"id":71},{"start":{"row":30,"column":11},"end":{"row":30,"column":17},"action":"insert","lines":["record"]}],[{"start":{"row":30,"column":17},"end":{"row":30,"column":18},"action":"insert","lines":["."],"id":72}],[{"start":{"row":30,"column":18},"end":{"row":30,"column":22},"action":"insert","lines":["Bins"],"id":73}],[{"start":{"row":29,"column":0},"end":{"row":29,"column":1},"action":"remove","lines":["\t"],"id":74,"ignore":true}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"insert","lines":[","],"id":75}],[{"start":{"row":12,"column":16},"end":{"row":12,"column":17},"action":"insert","lines":[" "],"id":76}],[{"start":{"row":12,"column":16},"end":{"row":12,"column":17},"action":"remove","lines":[" "],"id":77},{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"remove","lines":[","]}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"insert","lines":[","],"id":78}],[{"start":{"row":12,"column":16},"end":{"row":12,"column":17},"action":"insert","lines":[" "],"id":79}],[{"start":{"row":12,"column":16},"end":{"row":12,"column":17},"action":"remove","lines":[" "],"id":80},{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"remove","lines":[","]}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"insert","lines":[","],"id":81}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"remove","lines":[","],"id":82}],[{"start":{"row":11,"column":16},"end":{"row":12,"column":0},"action":"insert","lines":["",""],"id":83},{"start":{"row":12,"column":0},"end":{"row":12,"column":2},"action":"insert","lines":["\t\t"]},{"start":{"row":12,"column":2},"end":{"row":12,"column":3},"action":"insert","lines":["l"]},{"start":{"row":12,"column":3},"end":{"row":12,"column":4},"action":"insert","lines":["o"]},{"start":{"row":12,"column":4},"end":{"row":12,"column":5},"action":"insert","lines":["g"]},{"start":{"row":12,"column":5},"end":{"row":12,"column":6},"action":"insert","lines":["."]}],[{"start":{"row":12,"column":6},"end":{"row":12,"column":13},"action":"insert","lines":["Fatal()"],"id":84}],[{"start":{"row":12,"column":12},"end":{"row":12,"column":14},"action":"insert","lines":["\"\""],"id":85}],[{"start":{"row":12,"column":13},"end":{"row":12,"column":14},"action":"insert","lines":["o"],"id":86},{"start":{"row":12,"column":14},"end":{"row":12,"column":15},"action":"insert","lines":["k"]}],[{"start":{"row":12,"column":16},"end":{"row":12,"column":17},"action":"insert","lines":[")"],"id":87}],[{"start":{"row":12,"column":16},"end":{"row":12,"column":17},"action":"remove","lines":[")"],"id":88}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"insert","lines":["1"],"id":89}],[{"start":{"row":21,"column":16},"end":{"row":22,"column":0},"action":"insert","lines":["",""],"id":90},{"start":{"row":22,"column":0},"end":{"row":22,"column":2},"action":"insert","lines":["\t\t"]}],[{"start":{"row":22,"column":2},"end":{"row":23,"column":18},"action":"insert","lines":["","\t\tlog.Fatal(\"ok1\")"],"id":91}],[{"start":{"row":23,"column":15},"end":{"row":23,"column":16},"action":"remove","lines":["1"],"id":92}],[{"start":{"row":23,"column":15},"end":{"row":23,"column":16},"action":"insert","lines":["2"],"id":93}],[{"start":{"row":29,"column":16},"end":{"row":30,"column":0},"action":"insert","lines":["",""],"id":94},{"start":{"row":30,"column":0},"end":{"row":30,"column":2},"action":"insert","lines":["\t\t"]}],[{"start":{"row":30,"column":2},"end":{"row":31,"column":18},"action":"insert","lines":["","\t\tlog.Fatal(\"ok1\")"],"id":95}],[{"start":{"row":31,"column":15},"end":{"row":31,"column":16},"action":"remove","lines":["1"],"id":96}],[{"start":{"row":31,"column":15},"end":{"row":31,"column":16},"action":"insert","lines":["3"],"id":97}],[{"start":{"row":22,"column":0},"end":{"row":22,"column":2},"action":"remove","lines":["\t\t"],"id":98,"ignore":true},{"start":{"row":30,"column":0},"end":{"row":30,"column":2},"action":"remove","lines":["\t\t"]}],[{"start":{"row":22,"column":0},"end":{"row":23,"column":18},"action":"remove","lines":["","\t\tlog.Fatal(\"ok2\")"],"id":99},{"start":{"row":21,"column":16},"end":{"row":22,"column":0},"action":"remove","lines":["",""]}],[{"start":{"row":22,"column":12},"end":{"row":22,"column":13},"action":"insert","lines":["o"],"id":100},{"start":{"row":22,"column":13},"end":{"row":22,"column":14},"action":"insert","lines":["k"]}],[{"start":{"row":22,"column":13},"end":{"row":22,"column":14},"action":"remove","lines":["k"],"id":101}],[{"start":{"row":22,"column":13},"end":{"row":22,"column":14},"action":"insert","lines":["\""],"id":102}],[{"start":{"row":22,"column":13},"end":{"row":22,"column":14},"action":"remove","lines":["\""],"id":103},{"start":{"row":22,"column":12},"end":{"row":22,"column":13},"action":"remove","lines":["o"]}],[{"start":{"row":22,"column":12},"end":{"row":22,"column":13},"action":"insert","lines":["\""],"id":104},{"start":{"row":22,"column":13},"end":{"row":22,"column":14},"action":"insert","lines":["o"]},{"start":{"row":22,"column":14},"end":{"row":22,"column":15},"action":"insert","lines":["k"]},{"start":{"row":22,"column":15},"end":{"row":22,"column":16},"action":"insert","lines":["1"]},{"start":{"row":22,"column":16},"end":{"row":22,"column":17},"action":"insert","lines":["\""]}],[{"start":{"row":22,"column":17},"end":{"row":22,"column":18},"action":"insert","lines":[" "],"id":105},{"start":{"row":22,"column":18},"end":{"row":22,"column":19},"action":"insert","lines":["+"]}],[{"start":{"row":22,"column":19},"end":{"row":22,"column":20},"action":"insert","lines":[" "],"id":106}],[{"start":{"row":22,"column":23},"end":{"row":22,"column":24},"action":"insert","lines":["."],"id":107}],[{"start":{"row":22,"column":24},"end":{"row":22,"column":31},"action":"insert","lines":["Error()"],"id":108}],[{"start":{"row":11,"column":16},"end":{"row":13,"column":16},"action":"remove","lines":["","\t\tlog.Fatal(\"ok1\")","\t\tlog.Fatal(err)"],"id":109},{"start":{"row":11,"column":16},"end":{"row":12,"column":32},"action":"insert","lines":["","\t\tlog.Fatal(\"ok1\" + err.Error())"]}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"remove","lines":["1"],"id":110}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"insert","lines":["2"],"id":111}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"remove","lines":["2"],"id":112}],[{"start":{"row":12,"column":15},"end":{"row":12,"column":16},"action":"insert","lines":["1"],"id":113}],[{"start":{"row":21,"column":15},"end":{"row":21,"column":16},"action":"remove","lines":["1"],"id":114}],[{"start":{"row":21,"column":15},"end":{"row":21,"column":16},"action":"insert","lines":["2"],"id":115}],[{"start":{"row":27,"column":0},"end":{"row":29,"column":16},"action":"remove","lines":["","\t\tlog.Fatal(\"ok3\")","\t\tlog.Fatal(err)"],"id":116},{"start":{"row":27,"column":0},"end":{"row":28,"column":32},"action":"insert","lines":["","\t\tlog.Fatal(\"ok1\" + err.Error())"]}],[{"start":{"row":28,"column":15},"end":{"row":28,"column":16},"action":"remove","lines":["1"],"id":117}],[{"start":{"row":28,"column":15},"end":{"row":28,"column":16},"action":"insert","lines":["3"],"id":118}],[{"start":{"row":26,"column":16},"end":{"row":27,"column":0},"action":"remove","lines":["",""],"id":119}],[{"start":{"row":8,"column":48},"end":{"row":9,"column":0},"action":"insert","lines":["",""],"id":120},{"start":{"row":9,"column":0},"end":{"row":9,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":9,"column":1},"end":{"row":9,"column":2},"action":"insert","lines":["i"]},{"start":{"row":9,"column":2},"end":{"row":9,"column":3},"action":"insert","lines":["f"]}],[{"start":{"row":9,"column":3},"end":{"row":9,"column":4},"action":"insert","lines":[" "],"id":121}],[{"start":{"row":9,"column":4},"end":{"row":9,"column":6},"action":"insert","lines":["()"],"id":122}],[{"start":{"row":9,"column":5},"end":{"row":9,"column":6},"action":"insert","lines":["e"],"id":123},{"start":{"row":9,"column":6},"end":{"row":9,"column":7},"action":"insert","lines":["r"]}],[{"start":{"row":9,"column":6},"end":{"row":9,"column":7},"action":"remove","lines":["r"],"id":124},{"start":{"row":9,"column":5},"end":{"row":9,"column":6},"action":"remove","lines":["e"]},{"start":{"row":9,"column":4},"end":{"row":9,"column":6},"action":"remove","lines":["()"]}],[{"start":{"row":9,"column":4},"end":{"row":9,"column":5},"action":"insert","lines":["e"],"id":125},{"start":{"row":9,"column":5},"end":{"row":9,"column":6},"action":"insert","lines":["r"]},{"start":{"row":9,"column":6},"end":{"row":9,"column":7},"action":"insert","lines":["r"]}],[{"start":{"row":9,"column":7},"end":{"row":9,"column":8},"action":"insert","lines":[" "],"id":126},{"start":{"row":9,"column":8},"end":{"row":9,"column":9},"action":"insert","lines":["{"]}],[{"start":{"row":9,"column":9},"end":{"row":11,"column":2},"action":"insert","lines":["","\t\t","\t}"],"id":127}],[{"start":{"row":10,"column":2},"end":{"row":11,"column":32},"action":"insert","lines":["","\t\tlog.Fatal(\"ok1\" + err.Error())"],"id":128}],[{"start":{"row":11,"column":15},"end":{"row":11,"column":16},"action":"remove","lines":["1"],"id":129}],[{"start":{"row":11,"column":15},"end":{"row":11,"column":16},"action":"insert","lines":["0"],"id":130}],[{"start":{"row":10,"column":1},"end":{"row":10,"column":2},"action":"remove","lines":["\t"],"id":131},{"start":{"row":10,"column":0},"end":{"row":10,"column":1},"action":"remove","lines":["\t"]},{"start":{"row":9,"column":9},"end":{"row":10,"column":0},"action":"remove","lines":["",""]}],[{"start":{"row":9,"column":7},"end":{"row":9,"column":8},"action":"insert","lines":[" "],"id":132},{"start":{"row":9,"column":8},"end":{"row":9,"column":9},"action":"insert","lines":["!"]},{"start":{"row":9,"column":9},"end":{"row":9,"column":10},"action":"insert","lines":["="]}],[{"start":{"row":9,"column":10},"end":{"row":9,"column":11},"action":"insert","lines":[" "],"id":133},{"start":{"row":9,"column":11},"end":{"row":9,"column":12},"action":"insert","lines":["n"]},{"start":{"row":9,"column":12},"end":{"row":9,"column":13},"action":"insert","lines":["i"]},{"start":{"row":9,"column":13},"end":{"row":9,"column":14},"action":"insert","lines":["l"]}],[{"start":{"row":12,"column":24},"end":{"row":12,"column":33},"action":"remove","lines":["namespace"],"id":134},{"start":{"row":12,"column":24},"end":{"row":12,"column":25},"action":"insert","lines":["t"]},{"start":{"row":12,"column":25},"end":{"row":12,"column":26},"action":"insert","lines":["e"]},{"start":{"row":12,"column":26},"end":{"row":12,"column":27},"action":"insert","lines":["s"]},{"start":{"row":12,"column":27},"end":{"row":12,"column":28},"action":"insert","lines":["t"]}]]},"ace":{"folds":[],"scrolltop":0,"scrollleft":0,"selection":{"start":{"row":21,"column":18},"end":{"row":21,"column":18},"isBackwards":false},"options":{"guessTabSize":true,"useWrapMode":false,"wrapToView":true},"firstLineState":0},"timestamp":1602820172183,"hash":"0a9b052107c098b8c4e12a10e8864ec38f5b994a"}