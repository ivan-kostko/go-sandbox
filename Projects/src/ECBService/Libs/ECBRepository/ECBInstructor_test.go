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


package ECBRepository

import (
    "testing"
    ds "github.com/ivan-kostko/GoLibs/Repository/DataSource"
    repo "github.com/ivan-kostko/GoLibs/Repository"
    . "github.com/ivan-kostko/GoLibs/CustomErrors"
)

// Verify if ECBInstructor conforms repo.Instructor interface
var _ repo.Instructor = ECBInstructor{}

func TestECBInstructorGenerateInstruction(t *testing.T){

    ecbInstructor := ECBInstructor{}

    _ = []repo.FilteringCondition{

                    repo.FilteringCondition{
                        "FREQ":                 []repo.Predicate{ {repo.IN, []interface{}{"D"}}},
        //                "CURRENCY"            : repo.Predicate{ repo.IN, interface{}("USD")},
        //                "CURRENCY_DENOM"      : repo.Predicate{ repo.IN, interface{}("EUR")},
        //                "EXR_TYPE"            : repo.Predicate{ repo.IN, interface{}("SP00")},
        //                "EXR_SUFFIX"          : repo.Predicate{ repo.IN, interface{}("A")},
                    },

            }

    testCases := []struct{
        FilteringConditions []repo.FilteringCondition
        ExpectedInstruction ds.Instruction
        ExpectedError       *Error
    }{
        {
           []repo.FilteringCondition{

                    repo.FilteringCondition{
                        "FREQ"                : []repo.Predicate{ {repo.IN, []interface{}{"D"}}},
                        "CURRENCY"            : []repo.Predicate{ {repo.IN, []interface{}{"USD"}}},
                        "CURRENCY_DENOM"      : []repo.Predicate{ {repo.IN, []interface{}{"EUR"}}},
                        "EXR_TYPE"            : []repo.Predicate{ {repo.IN, []interface{}{"SP00"}}},
                        "EXR_SUFFIX"          : []repo.Predicate{ {repo.IN, []interface{}{"A"}}},
                    },

            },
            ds.Instruction("D.USD.EUR.SP00.A"),
            nil,
        },

    }

    for _, testCase := range testCases {
        actualInstruction, actualError := ecbInstructor.GenerateInstruction(testCase.FilteringConditions...)
        if actualInstruction != testCase.ExpectedInstruction ||
            actualError != testCase.ExpectedError {
                t.Errorf("ecbInstructor.GenerateInstruction(%#v) returned: \r\n Instruction: %#v and Error: %#v \r\n while expected Instruction: %#v and Error: %#v", testCase.FilteringConditions, actualInstruction, actualError, testCase.ExpectedInstruction, testCase.ExpectedError)
            }
    }

}
