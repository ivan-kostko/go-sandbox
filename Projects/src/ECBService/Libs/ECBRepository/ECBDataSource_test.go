//   Copyright (c) 2016 Ivan A Kostko (github.com/ivan-kostko)

//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

//       http://www.apache.org/licenses/LICENSE-2.0

//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.


// ECBRepository project ECBRepository.go
package ECBRepository

import (
    "testing"

)

var dsConf = ECBDataSourceConfiguration{
            "https://a-sdw-wsrest.ecb.int/service/data/EXR",
            map[string][]string{"Accept":{"application/vnd.sdmx.data+json;version=1.0.0-wd"}},
            nil,

}

func TestECBDataSource(t *testing.T){
    ds, err := getECBDataSource(dsConf)
    if err != nil {
        t.Fatalf("getECBDataSource returned error as %v ",err)
    }
    result, err := ds.ExecuteInstruction("D.USD.EUR.SP00.A?startPeriod=2016-05-05&endPeriod=2016-05-16")
    if err != nil {
        t.Errorf("ds.ExecuteInstruction returned error as %#v ",err)
    }
    t.Logf("ds.ExecuteInstruction returned result as %s ", result)
}
