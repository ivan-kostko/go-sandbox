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
    repo "github.com/ivan-kostko/GoLibs/Repository"
    //business "ECBService/Libs/Models/Business"
)


// Represents the ECB Currency Exchange rate repository implementation
type ECBCurrencyExchangeRepo struct {
    repo *repo.Repository
}


// Represents configuration for ECB Currency Exchange rate repository implementation
type ECBCurrencyExchangeRepoConfig struct{
    DSConf       ECBDataSourceConfiguration
}

// ECBCurrencyExchangeRepo factory
func GetNewECBCurrencyExchangeRepo(repoConfig ECBCurrencyExchangeRepoConfig) *ECBCurrencyExchangeRepo {
    ecbDS , _ := getECBDataSource(repoConfig.DSConf)

    return &ECBCurrencyExchangeRepo{repo: repo.GetNewRepository(ecbDS,nil,nil,nil)}
}


