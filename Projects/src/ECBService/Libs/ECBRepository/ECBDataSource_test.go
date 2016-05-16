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

func TestECBDataSource(t *testing.T){
    ds := getECBDataSource()
    result, err := ds.ExecuteInstruction("https://sdw-wsrest.ecb.europa.eu/service/data/EXR/D.USD.EUR.SP00.A?startPeriod=2016-05-05&endPeriod=2016-05-16")
    if err != nil {
        t.Errorf("ds.ExecuteInstruction returned error as %#v ",err)
    }
    t.Logf("ds.ExecuteInstruction returned result as %s ", result)
}
