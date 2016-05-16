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


// Default project Default.go
package Default

import (
    ecb "ECBService/Libs/Models/ECB"
)

var registerAs = "sdw-wsrest.ecb.europa.eu/Response/MessageHeader/Default"

func init() {
    ecb.RegisterFactory(registerAs, GetNewMessageHeader)
}

type MessageHeader struct{
    //XMLName  xml.Name  `xml:'message:GenericData'`
    ID             string    `xml:"Header>ID"`
    Sender         struct{
                       Id    string `xml:"id,attr"`
                   }    `xml:"Header>Sender"`
    Structure      struct{
                        StructureID            string `xml:"structureID,attr"`
                        DimensionAtObservation string `xml:"dimensionAtObservation,attr"`
                        URN                    string `xml:"Structure>URN"`
                    }    `xml:"Header>Structure"`
}

func GetNewMessageHeader() interface{} {
    return MessageHeader{}
}
