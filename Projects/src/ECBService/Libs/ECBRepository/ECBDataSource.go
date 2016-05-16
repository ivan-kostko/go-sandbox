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

import(
    "fmt"
    "io/ioutil"
	"net/http"
    ds "ECBService/Libs/Repository/DataSource"

    . "github.com/ivan-kostko/GoLibs/CustomErrors"
)



func getECBDataSource() *ds.DataSource {

    client := http.DefaultClient
    header := http.Header{"Accept":{"application/vnd.sdmx.data+json;version=1.0.0-wd"}}

    executeInstruction := ds.ExecuteInstruction(func(i ds.Instruction) (ds.Result, *Error) {
            url := string(i)
            request, err := http.NewRequest("GET", url, nil)
            if err != nil {
                NewError(InvalidOperation, err.Error())
            }
            request.Header = header
            res, err := client.Do(request)
        	if err != nil {
        		fmt.Println("Error on http.Get(): ",err)
        	}
            result, err := ioutil.ReadAll(res.Body)
        	res.Body.Close()
        	if err != nil {
        		fmt.Println("Error on ioutil.ReadAll(res.Body): ",err)
        	}
            return result, nil
    })

    return &ds.DataSource{
        ExecuteInstruction: executeInstruction,
    }
}
