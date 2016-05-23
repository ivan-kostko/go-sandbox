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

import(
    ds "github.com/ivan-kostko/GoLibs/Repository/DataSource"
    repo "github.com/ivan-kostko/GoLibs/Repository"
    . "github.com/ivan-kostko/GoLibs/CustomErrors"

    "strings"
)

const(
    ERR_MORETHANONEFILTERINGCONDITION  = "ECBInstructor: Currenly, disjuction functionality is not supported even by "
    ERR_NONSUPORTEDOPERATOR            = "ECBInstructor: Currenly, the operator is not supported"
)

const (
    FREQ            =  "FREQ"
    CURRENCY        =  "CURRENCY"
    CURRENCY_DENOM  =  "CURRENCY_DENOM"
    EXR_TYPE        =  "EXR_TYPE"
    EXR_SUFFIX      =  "EXR_SUFFIX"


)

type ECBInstructor struct{
    keyPathParamOR            string
    keyPathParamDelimeter     string

}



// Generates Instruction for ECB data source
// Returns Instruction(""),
func (this ECBInstructor) GenerateInstruction(fcs ...repo.FilteringCondition) (ds.Instruction, *Error){
    if len(fc) > 1 {
        return ds.Instruction(""), NewError(Nonsupported, ERR_MORETHANONEFILTERINGCONDITION)
    }
    fc := fcs[0]

    frquencies := ""
    for frquncy := range fc[FREQ]
    return ds.Instruction(""), nil
}

func (this ECBInstructor) generateKeyPathParameter(fc repo.FilteringCondition) (string, *Error){
    frequencyKey := this.generateFrequency(fc[FREQ])

}


func (this ECBInstructor) generateFrequency(freqPredicates []repo.Predicate) (string, *Error){
    ret := ""
    for i, frequncy := range freqPredictes {
        if frequncy.Operator == repo.IN {
            ret = strings.Join({ret, frequncy.Values}, this.keyPathParamOR)

        }else {
            return ds.Instruction(""), NewError(Nonsupported, ERR_NONSUPORTEDOPERATOR)
        }
    }
    return ret, nil
}


