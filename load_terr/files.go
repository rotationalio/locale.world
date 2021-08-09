package main

import "encoding/xml"

type Territories struct {
	XMLName     xml.Name    `xml:"territoryInfo"`
	Territories []Territory `xml:"territory"`
}

type Territory struct {
	Code                string              `xml:"type,attr"`
	Gdp                 string              `xml:"gdp,attr"`
	LiteracyPct         float32             `xml:"literacyPercent,attr"`
	Population          int                 `xml:"population,attr"`
	LanguageTerritories []LanguageTerritory `xml:"languagePopulation"`
}

type LanguageTerritory struct {
	Code           string `xml:"type,attr"`
	PopulationPct  string `xml:"populationPercent,attr"`
	OfficialStatus string `xml:"officialStatus,attr"`
}

type TerrSubs struct {
	TerrSubs []TerrSub `json:"territorySubdivision"`
}

type TerrSub struct {
	TerrCode     string        `json:"terrCode"`
	TerrName     string        `json:"terrName"`
	Subdivisions []Subdivision `json:"subdivisions"`
}

type Subdivision struct {
	Code string `json:"subdivCode"`
	Name string `json:"subdivName"`
}

type Languages struct {
	Languages []Language `json:"languages"`
}

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CIACountries struct {
	Territories []CIACountry `json:"Territories"`
}

type CIACountry struct {
	Name                                                 string `json:"Name"`
	Background                                           string `json:"Background"`
	Location                                             string `json:"Location"`
	Geographic_coordinates                               string `json:"Geographic_coordinates"`
	Map_references                                       string `json:"Map_references"`
	Area                                                 string `json:"Area"`
	Area_comparative                                     string `json:"Area_comparative"`
	Land_boundaries                                      string `json:"Land_boundaries"`
	Coastline                                            string `json:"Coastline"`
	Maritime_claims                                      string `json:"Maritime_claims"`
	Climate                                              string `json:"Climate"`
	Terrain                                              string `json:"Terrain"`
	Elevation                                            string `json:"Elevation"`
	Natural_resources                                    string `json:"Natural_resources"`
	Land_use                                             string `json:"Land_use"`
	Irrigated_land                                       string `json:"Irrigated_land"`
	Total_renewable_water_resources                      string `json:"Total_renewable_water_resources"`
	Population_distribution                              string `json:"Population_distribution"`
	Demographic_profile                                  string `json:"Demographic_profile"`
	Railways                                             string `json:"Railways"`
	Natural_hazards                                      string `json:"Natural_hazards"`
	Environment_international_agreements                 string `json:"Environment_international_agreements"`
	Geography_note                                       string `json:"Geography_note"`
	Nationality                                          string `json:"Nationality"`
	Ethnic_groups                                        string `json:"Ethnic_groups"`
	Languages                                            string `json:"Languages"`
	Religions                                            string `json:"Religions"`
	Age_structure                                        string `json:"Age_structure"`
	Dependency_ratios                                    string `json:"Dependency_ratios"`
	Median_age                                           string `json:"Median_age"`
	Population_growth_rate                               string `json:"Population_growth_rate"`
	Birth_rate                                           string `json:"Birth_rate"`
	Death_rate                                           string `json:"Death_rate"`
	Net_migration_rate                                   string `json:"Net_migration_rate"`
	Urbanization                                         string `json:"Urbanization"`
	Major_urban_areas_population                         string `json:"Major_urban_areas_population"`
	Sex_ratio                                            string `json:"Sex_ratio"`
	Mothers_mean_age_at_first_birth                      string `json:"Mothers_mean_age_at_first_birth"`
	Maternal_mortality_rate                              string `json:"Maternal_mortality_rate"`
	Infant_mortality_rate                                string `json:"Infant_mortality_rate"`
	Life_expectancy_at_birth                             string `json:"Life_expectancy_at_birth"`
	Total_fertility_rate                                 string `json:"Total_fertility_rate"`
	Contraceptive_prevalence_rate                        string `json:"Contraceptive_prevalence_rate"`
	Drinking_water_source                                string `json:"Drinking_water_source"`
	Current_Health_Expenditure                           string `json:"Current_Health_Expenditure"`
	Physicians_density                                   string `json:"Physicians_density"`
	Hospital_bed_density                                 string `json:"Hospital_bed_density"`
	Sanitation_facility_access                           string `json:"Sanitation_facility_access"`
	HIV_AIDS_adult_prevalence_rate                       string `json:"HIV_AIDS_adult_prevalence_rate"`
	HIV_AIDS_people_living_with_HIV_AIDS                 string `json:"HIV_AIDS_people_living_with_HIV_AIDS"`
	HIV_AIDS_deaths                                      string `json:"HIV_AIDS_deaths"`
	Major_infectious_diseases                            string `json:"Major_infectious_diseases"`
	Obesity_adult_prevalence_rate                        string `json:"Obesity_adult_prevalence_rate"`
	Children_under_the_age_of__years_underweight         string `json:"Children_under_the_age_of__years_underweight"`
	Education_expenditures                               string `json:"Education_expenditures"`
	Literacy                                             string `json:"Literacy"`
	School_life_expectancy_primary_to_tertiary_education string `json:"School_life_expectancy_primary_to_tertiary_education"`
	Unemployment_youth_ages_                             string `json:"Unemployment_youth_ages_"`
	Environment_current_issues                           string `json:"Environment_current_issues"`
	Air_pollutants                                       string `json:"Air_pollutants"`
	Total_water_withdrawal                               string `json:"Total_water_withdrawal"`
	Revenue_from_forest_resources                        string `json:"Revenue_from_forest_resources"`
	Revenue_from_coal                                    string `json:"Revenue_from_coal"`
	Food_insecurity                                      string `json:"Food_insecurity"`
	Waste_and_recycling                                  string `json:"Waste_and_recycling"`
	Country_name                                         string `json:"Country_name"`
	Government_type                                      string `json:"Government_type"`
	Capital                                              string `json:"Capital"`
	Administrative_divisions                             string `json:"Administrative_divisions"`
	Independence                                         string `json:"Independence"`
	National_holiday                                     string `json:"National_holiday"`
	Constitution                                         string `json:"Constitution"`
	Legal_system                                         string `json:"Legal_system"`
	International_law_organization_participation         string `json:"International_law_organization_participation"`
	Citizenship                                          string `json:"Citizenship"`
	Suffrage                                             string `json:"Suffrage"`
	Executive_branch                                     string `json:"Executive_branch"`
	Legislative_branch                                   string `json:"Legislative_branch"`
	Judicial_branch                                      string `json:"Judicial_branch"`
	Political_parties_and_leaders                        string `json:"Political_parties_and_leaders"`
	International_organization_participation             string `json:"International_organization_participation"`
	Diplomatic_representation_in_the_US                  string `json:"Diplomatic_representation_in_the_US"`
	Diplomatic_representation_from_the_US                string `json:"Diplomatic_representation_from_the_US"`
	Flag_description                                     string `json:"Flag_description"`
	National_symbols                                     string `json:"National_symbols"`
	National_anthem                                      string `json:"National_anthem"`
	Economic_overview                                    string `json:"Economic_overview"`
	Real_GDP_growth_rate                                 string `json:"Real_GDP_growth_rate"`
	Inflation_rate_consumer_prices                       string `json:"Inflation_rate_consumer_prices"`
	Real_GDP_purchasing_power_parity                     string `json:"Real_GDP_purchasing_power_parity"`
	GDP_official_exchange_rate                           string `json:"GDP_official_exchange_rate"`
	Real_GDP_per_capita                                  string `json:"Real_GDP_per_capita"`
	Gross_national_saving                                string `json:"Gross_national_saving"`
	GDP_composition_by_sector_of_origin                  string `json:"GDP_composition_by_sector_of_origin"`
	GDP_composition_by_end_use                           string `json:"GDP_composition_by_end_use"`
	Ease_of_Doing_Business_Index_scores                  string `json:"Ease_of_Doing_Business_Index_scores"`
	Agricultural_products                                string `json:"Agricultural_products"`
	Industries                                           string `json:"Industries"`
	Industrial_production_growth_rate                    string `json:"Industrial_production_growth_rate"`
	Labor_force                                          string `json:"Labor_force"`
	Labor_force_by_occupation                            string `json:"Labor_force_by_occupation"`
	Unemployment_rate                                    string `json:"Unemployment_rate"`
	Population_below_poverty_line                        string `json:"Population_below_poverty_line"`
	Gini_Index_coefficient_distribution_of_family_income string `json:"Gini_Index_coefficient_distribution_of_family_income"`
	Household_income_or_consumption_by_percentage_share  string `json:"Household_income_or_consumption_by_percentage_share"`
	Budget                                               string `json:"Budget"`
	Taxes_and_other_revenues                             string `json:"Taxes_and_other_revenues"`
	Budget_surplus__or_deficit_                          string `json:"Budget_surplus__or_deficit_"`
	Public_debt                                          string `json:"Public_debt"`
	Fiscal_year                                          string `json:"Fiscal_year"`
	Current_account_balance                              string `json:"Current_account_balance"`
	Exports                                              string `json:"Exports"`
	Exports_partners                                     string `json:"Exports_partners"`
	Exports_commodities                                  string `json:"Exports_commodities"`
	Imports                                              string `json:"Imports"`
	Imports_partners                                     string `json:"Imports_partners"`
	Imports_commodities                                  string `json:"Imports_commodities"`
	Reserves_of_foreign_exchange_and_gold                string `json:"Reserves_of_foreign_exchange_and_gold"`
	Debt_external                                        string `json:"Debt_external"`
	Exchange_rates                                       string `json:"Exchange_rates"`
	Electricity_access                                   string `json:"Electricity_access"`
	Electricity_production                               string `json:"Electricity_production"`
	Electricity_consumption                              string `json:"Electricity_consumption"`
	Electricity_exports                                  string `json:"Electricity_exports"`
	Electricity_imports                                  string `json:"Electricity_imports"`
	Electricity_installed_generating_capacity            string `json:"Electricity_installed_generating_capacity"`
	Electricity_from_fossil_fuels                        string `json:"Electricity_from_fossil_fuels"`
	Electricity_from_nuclear_fuels                       string `json:"Electricity_from_nuclear_fuels"`
	Electricity_from_hydroelectric_plants                string `json:"Electricity_from_hydroelectric_plants"`
	Electricity_from_other_renewable_sources             string `json:"Electricity_from_other_renewable_sources"`
	Crude_oil_production                                 string `json:"Crude_oil_production"`
	Crude_oil_exports                                    string `json:"Crude_oil_exports"`
	Crude_oil_imports                                    string `json:"Crude_oil_imports"`
	Crude_oil_proved_reserves                            string `json:"Crude_oil_proved_reserves"`
	Refined_petroleum_products_production                string `json:"Refined_petroleum_products_production"`
	Refined_petroleum_products_consumption               string `json:"Refined_petroleum_products_consumption"`
	Refined_petroleum_products_exports                   string `json:"Refined_petroleum_products_exports"`
	Refined_petroleum_products_imports                   string `json:"Refined_petroleum_products_imports"`
	Natural_gas_production                               string `json:"Natural_gas_production"`
	Natural_gas_consumption                              string `json:"Natural_gas_consumption"`
	Natural_gas_exports                                  string `json:"Natural_gas_exports"`
	Natural_gas_imports                                  string `json:"Natural_gas_imports"`
	Natural_gas_proved_reserves                          string `json:"Natural_gas_proved_reserves"`
	Carbon_dioxide_emissions_from_consumption_of_energy  string `json:"Carbon_dioxide_emissions_from_consumption_of_energy"`
	Telephones_fixed_lines                               string `json:"Telephones_fixed_lines"`
	Telephones_mobile_cellular                           string `json:"Telephones_mobile_cellular"`
	Telecommunication_systems                            string `json:"Telecommunication_systems"`
	Broadcast_media                                      string `json:"Broadcast_media"`
	Internet_country_code                                string `json:"Internet_country_code"`
	Internet_users                                       string `json:"Internet_users"`
	Broadband_fixed_subscriptions                        string `json:"Broadband_fixed_subscriptions"`
	National_air_transport_system                        string `json:"National_air_transport_system"`
	Civil_aircraft_registration_country_code_prefix      string `json:"Civil_aircraft_registration_country_code_prefix"`
	Airports                                             string `json:"Airports"`
	Airports_with_paved_runways                          string `json:"Airports_with_paved_runways"`
	Airports_with_unpaved_runways                        string `json:"Airports_with_unpaved_runways"`
	Heliports                                            string `json:"Heliports"`
	Pipelines                                            string `json:"Pipelines"`
	Roadways                                             string `json:"Roadways"`
	Waterways                                            string `json:"Waterways"`
	Ports_and_terminals                                  string `json:"Ports_and_terminals"`
	Military_and_security_forces                         string `json:"Military_and_security_forces"`
	Military_expenditures                                string `json:"Military_expenditures"`
	Military_and_security_service_personnel_strengths    string `json:"Military_and_security_service_personnel_strengths"`
	Military_equipment_inventories_and_acquisitions      string `json:"Military_equipment_inventories_and_acquisitions"`
	Military_service_age_and_obligation                  string `json:"Military_service_age_and_obligation"`
	Military_note                                        string `json:"Military_note"`
	Terrorist_groups                                     string `json:"Terrorist_groups"`
	Disputes_international                               string `json:"Disputes_international"`
	Refugees_and_internally_displaced_persons            string `json:"Refugees_and_internally_displaced_persons"`
	Illicit_drugs                                        string `json:"Illicit_drugs"`
	Trafficking_in_persons                               string `json:"Trafficking_in_persons"`
	Timezone                                             string `json:"Timezone"`
	Gallery                                              string `json:"Gallery"`
	Flag                                                 string `json:"Flag"`
	Map                                                  string `json:"Map"`
}
