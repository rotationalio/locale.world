package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "19910606"
	dbname   = "locale"
)

func main() {
	xmlFile, err := os.Open("../data/territories.xml")
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var territories Territories
	xml.Unmarshal(byteValue, &territories)

	jsonFile, err := os.Open("../data/subdivisions.json")
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	var terrSubs TerrSubs
	json.Unmarshal(byteValue, &terrSubs)

	jsonFile, err = os.Open("../data/languages.json")
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	var languages Languages
	json.Unmarshal(byteValue, &languages)

	jsonFile, err = os.Open("../data/cia_countries.json")
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	var ciaCountries CIACountries
	json.Unmarshal(byteValue, &ciaCountries)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	if _, err := db.Exec(terrTableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(langTableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(langTerrTableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(subTableCreationQuery); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(languages.Languages); i++ {
		_, e := db.Exec(`
			INSERT INTO languages (code, name) VALUES ($1,$2)
			ON CONFLICT DO NOTHING`,
			languages.Languages[i].Code, languages.Languages[i].Name)
		CheckError(e)
	}

	for i := 0; i < len(territories.Territories); i++ {
		languageTerritories := territories.Territories[i].LanguageTerritories

		sqlStatement := `
		INSERT INTO territories (code, gdp, literacy_pct, population, main_language)
		VALUES ($1,$2,$3,$4,$5)`
		_, e := db.Exec(sqlStatement,
			territories.Territories[i].Code, territories.Territories[i].Gdp,
			territories.Territories[i].LiteracyPct, territories.Territories[i].Population,
			languageTerritories[0].Code)
		CheckError(e)

		var officialStatus bool
		for j := 0; j < len(languageTerritories); j++ {
			if languageTerritories[j].OfficialStatus == "official" {
				officialStatus = true
			} else {
				officialStatus = false
			}
			sqlStatement := `
			INSERT INTO languageterritory (language_code, territory_code, population_pct, official_status)
			VALUES ($1,$2,$3,$4)`
			_, e = db.Exec(sqlStatement,
				languageTerritories[j].Code, territories.Territories[i].Code,
				languageTerritories[j].PopulationPct, officialStatus)
			CheckError(e)

		}
	}

	for i := 0; i < len(terrSubs.TerrSubs); i++ {
		_, e := db.Exec(`
			UPDATE territories
			SET name = $1
			WHERE code = $2`,
			terrSubs.TerrSubs[i].TerrName, terrSubs.TerrSubs[i].TerrCode)
		CheckError(e)

		subdivs := terrSubs.TerrSubs[i].Subdivisions
		for j := 0; j < len(subdivs); j++ {
			if subdivs[j].Code != "" {
				_, e := db.Exec(`
					INSERT INTO subdivisions (code, name, territory_code) VALUES ($1,$2,$3)`,
					subdivs[j].Code, subdivs[j].Name, terrSubs.TerrSubs[i].TerrCode)
				CheckError(e)
			}
		}
	}

	for i := 0; i < len(ciaCountries.Territories); i++ {
		_, e := db.Exec(`
			UPDATE territories
			SET background = $1,
					timezone = $3,
					location = $4,
					geographic_coordinates = $5,
					map_references = $6,
					area = $7,
					area_comparative = $8,
					land_boundaries = $9,
					coastline = $10,
					maritime_claims = $11,
					climate = $12,
					terrain = $13,
					elevation = $14,
					natural_resources = $15,
					land_use = $16,
					irrigated_land = $17,
					total_renewable_water_resources = $18,
					population_distribution = $19,
					natural_hazards = $20,
					environment_international_agreements = $21,
					geography_note = $22,
					nationality = $23,
					ethnic_groups = $24,
					languages = $25,
					religions = $26,
					demographic_profile = $27,
					age_structure = $28,
					dependency_ratios = $29,
					median_age = $30,
					population_growth_rate = $31,
					birth_rate = $32,
					death_rate = $33,
					net_migration_rate = $34,
					urbanization = $35,
					major_urban_areas_population = $36,
					sex_ratio = $37,
					mothers_mean_age_at_first_birth = $38,
					maternal_mortality_rate = $39,
					infant_mortality_rate = $40,
					life_expectancy_at_birth = $41,
					total_fertility_rate = $42,
					contraceptive_prevalence_rate = $43,
					drinking_water_source = $44,
					current_health_expenditure = $45,
					physicians_density = $46,
					hospital_bed_density = $47,
					sanitation_facility_access = $48,
					hiv_aids_adult_prevalence_rate = $49,
					hiv_aids_people_living_with_hiv_aids = $50,
					hiv_aids_deaths = $51,
					major_infectious_diseases = $52,
					obesity_adult_prevalence_rate = $53,
					children_under_the_age_of__years_underweight = $54,
					education_expenditures = $55,
					literacy = $56,
					school_life_expectancy_primary_to_tertiary_education = $57,
					unemployment_youth_ages_ = $58,
					environment_current_issues = $59,
					air_pollutants = $60,
					total_water_withdrawal = $61,
					revenue_from_forest_resources = $62,
					revenue_from_coal = $63,
					food_insecurity = $64,
					waste_and_recycling = $65,
					country_name = $66,
					government_type = $67,
					capital = $68,
					administrative_divisions = $69,
					independence = $70,
					national_holiday = $71,
					constitution = $72,
					legal_system = $73,
					international_law_organization_participation = $74,
					citizenship = $75,
					suffrage = $76,
					executive_branch = $77,
					legislative_branch = $78,
					judicial_branch = $79,
					political_parties_and_leaders = $80,
					international_organization_participation = $81,
					diplomatic_representation_in_the_us = $82,
					diplomatic_representation_from_the_us = $83,
					flag_description = $84,
					national_symbols = $85,
					national_anthem = $86,
					economic_overview = $87,
					real_gdp_growth_rate = $88,
					inflation_rate_consumer_prices = $89,
					real_gdp_purchasing_power_parity = $90,
					gdp_official_exchange_rate = $91,
					real_gdp_per_capita = $92,
					gross_national_saving = $93,
					gdp_composition_by_sector_of_origin = $94,
					gdp_composition_by_end_use = $95,
					ease_of_doing_business_index_scores = $96,
					agricultural_products = $97,
					industries = $98,
					industrial_production_growth_rate = $99,
					labor_force = $100,
					labor_force_by_occupation = $101,
					unemployment_rate = $102,
					population_below_poverty_line = $103,
					gini_index_coefficient_distribution_of_family_income = $104,
					household_income_or_consumption_by_percentage_share = $105,
					budget = $106,
					taxes_and_other_revenues = $107,
					budget_surplus__or_deficit_ = $108,
					public_debt = $109,
					fiscal_year = $110,
					current_account_balance = $111,
					exports = $112,
					exports_partners = $113,
					exports_commodities = $114,
					imports = $115,
					imports_partners = $116,
					imports_commodities = $117,
					reserves_of_foreign_exchange_and_gold = $118,
					debt_external = $119,
					exchange_rates = $120,
					electricity_access = $121,
					electricity_production = $122,
					electricity_consumption = $123,
					electricity_exports = $124,
					electricity_imports = $125,
					electricity_installed_generating_capacity = $126,
					electricity_from_fossil_fuels = $127,
					electricity_from_nuclear_fuels = $128,
					electricity_from_hydroelectric_plants = $129,
					electricity_from_other_renewable_sources = $130,
					crude_oil_production = $131,
					crude_oil_exports = $132,
					crude_oil_imports = $133,
					crude_oil_proved_reserves = $134,
					refined_petroleum_products_production = $135,
					refined_petroleum_products_consumption = $136,
					refined_petroleum_products_exports = $137,
					refined_petroleum_products_imports = $138,
					natural_gas_production = $139,
					natural_gas_consumption = $140,
					natural_gas_exports = $141,
					natural_gas_imports = $142,
					natural_gas_proved_reserves = $143,
					carbon_dioxide_emissions_from_consumption_of_energy = $144,
					telephones_fixed_lines = $145,
					telephones_mobile_cellular = $146,
					telecommunication_systems = $147,
					broadcast_media = $148,
					internet_country_code = $149,
					internet_users = $150,
					broadband_fixed_subscriptions = $151,
					national_air_transport_system = $152,
					civil_aircraft_registration_country_code_prefix = $153,
					airports = $154,
					airports_with_paved_runways = $155,
					airports_with_unpaved_runways = $156,
					pipelines = $157,
					railways = $158,
					roadways = $159,
					waterways = $160,
					ports_and_terminals = $161,
					military_and_security_forces = $162,
					military_expenditures = $163,
					military_and_security_service_personnel_strengths = $164,
					military_equipment_inventories_and_acquisitions = $165,
					military_service_age_and_obligation = $166,
					military_note = $167,
					disputes_international = $168,
					refugees_and_internally_displaced_persons = $169,
					trafficking_in_persons = $170,
					illicit_drugs = $171,
					gallery = $172,
					flag = $173,
					map = $174
			WHERE name = $2`,
			ciaCountries.Territories[i].Background,
			ciaCountries.Territories[i].Name,
			ciaCountries.Territories[i].Timezone,
			ciaCountries.Territories[i].Location,
			ciaCountries.Territories[i].Geographic_coordinates,
			ciaCountries.Territories[i].Map_references,
			ciaCountries.Territories[i].Area,
			ciaCountries.Territories[i].Area_comparative,
			ciaCountries.Territories[i].Land_boundaries,
			ciaCountries.Territories[i].Coastline,
			ciaCountries.Territories[i].Maritime_claims,
			ciaCountries.Territories[i].Climate,
			ciaCountries.Territories[i].Terrain,
			ciaCountries.Territories[i].Elevation,
			ciaCountries.Territories[i].Natural_resources,
			ciaCountries.Territories[i].Land_use,
			ciaCountries.Territories[i].Irrigated_land,
			ciaCountries.Territories[i].Total_renewable_water_resources,
			ciaCountries.Territories[i].Population_distribution,
			ciaCountries.Territories[i].Natural_hazards,
			ciaCountries.Territories[i].Environment_international_agreements,
			ciaCountries.Territories[i].Geography_note,
			ciaCountries.Territories[i].Nationality,
			ciaCountries.Territories[i].Ethnic_groups,
			ciaCountries.Territories[i].Languages,
			ciaCountries.Territories[i].Religions,
			ciaCountries.Territories[i].Demographic_profile,
			ciaCountries.Territories[i].Age_structure,
			ciaCountries.Territories[i].Dependency_ratios,
			ciaCountries.Territories[i].Median_age,
			ciaCountries.Territories[i].Population_growth_rate,
			ciaCountries.Territories[i].Birth_rate,
			ciaCountries.Territories[i].Death_rate,
			ciaCountries.Territories[i].Net_migration_rate,
			ciaCountries.Territories[i].Urbanization,
			ciaCountries.Territories[i].Major_urban_areas_population,
			ciaCountries.Territories[i].Sex_ratio,
			ciaCountries.Territories[i].Mothers_mean_age_at_first_birth,
			ciaCountries.Territories[i].Maternal_mortality_rate,
			ciaCountries.Territories[i].Infant_mortality_rate,
			ciaCountries.Territories[i].Life_expectancy_at_birth,
			ciaCountries.Territories[i].Total_fertility_rate,
			ciaCountries.Territories[i].Contraceptive_prevalence_rate,
			ciaCountries.Territories[i].Drinking_water_source,
			ciaCountries.Territories[i].Current_Health_Expenditure,
			ciaCountries.Territories[i].Physicians_density,
			ciaCountries.Territories[i].Hospital_bed_density,
			ciaCountries.Territories[i].Sanitation_facility_access,
			ciaCountries.Territories[i].HIV_AIDS_adult_prevalence_rate,
			ciaCountries.Territories[i].HIV_AIDS_people_living_with_HIV_AIDS,
			ciaCountries.Territories[i].HIV_AIDS_deaths,
			ciaCountries.Territories[i].Major_infectious_diseases,
			ciaCountries.Territories[i].Obesity_adult_prevalence_rate,
			ciaCountries.Territories[i].Children_under_the_age_of__years_underweight,
			ciaCountries.Territories[i].Education_expenditures,
			ciaCountries.Territories[i].Literacy,
			ciaCountries.Territories[i].School_life_expectancy_primary_to_tertiary_education,
			ciaCountries.Territories[i].Unemployment_youth_ages_,
			ciaCountries.Territories[i].Environment_current_issues,
			ciaCountries.Territories[i].Air_pollutants,
			ciaCountries.Territories[i].Total_water_withdrawal,
			ciaCountries.Territories[i].Revenue_from_forest_resources,
			ciaCountries.Territories[i].Revenue_from_coal,
			ciaCountries.Territories[i].Food_insecurity,
			ciaCountries.Territories[i].Waste_and_recycling,
			ciaCountries.Territories[i].Country_name,
			ciaCountries.Territories[i].Government_type,
			ciaCountries.Territories[i].Capital,
			ciaCountries.Territories[i].Administrative_divisions,
			ciaCountries.Territories[i].Independence,
			ciaCountries.Territories[i].National_holiday,
			ciaCountries.Territories[i].Constitution,
			ciaCountries.Territories[i].Legal_system,
			ciaCountries.Territories[i].International_law_organization_participation,
			ciaCountries.Territories[i].Citizenship,
			ciaCountries.Territories[i].Suffrage,
			ciaCountries.Territories[i].Executive_branch,
			ciaCountries.Territories[i].Legislative_branch,
			ciaCountries.Territories[i].Judicial_branch,
			ciaCountries.Territories[i].Political_parties_and_leaders,
			ciaCountries.Territories[i].International_organization_participation,
			ciaCountries.Territories[i].Diplomatic_representation_in_the_US,
			ciaCountries.Territories[i].Diplomatic_representation_from_the_US,
			ciaCountries.Territories[i].Flag_description,
			ciaCountries.Territories[i].National_symbols,
			ciaCountries.Territories[i].National_anthem,
			ciaCountries.Territories[i].Economic_overview,
			ciaCountries.Territories[i].Real_GDP_growth_rate,
			ciaCountries.Territories[i].Inflation_rate_consumer_prices,
			ciaCountries.Territories[i].Real_GDP_purchasing_power_parity,
			ciaCountries.Territories[i].GDP_official_exchange_rate,
			ciaCountries.Territories[i].Real_GDP_per_capita,
			ciaCountries.Territories[i].Gross_national_saving,
			ciaCountries.Territories[i].GDP_composition_by_sector_of_origin,
			ciaCountries.Territories[i].GDP_composition_by_end_use,
			ciaCountries.Territories[i].Ease_of_Doing_Business_Index_scores,
			ciaCountries.Territories[i].Agricultural_products,
			ciaCountries.Territories[i].Industries,
			ciaCountries.Territories[i].Industrial_production_growth_rate,
			ciaCountries.Territories[i].Labor_force,
			ciaCountries.Territories[i].Labor_force_by_occupation,
			ciaCountries.Territories[i].Unemployment_rate,
			ciaCountries.Territories[i].Population_below_poverty_line,
			ciaCountries.Territories[i].Gini_Index_coefficient_distribution_of_family_income,
			ciaCountries.Territories[i].Household_income_or_consumption_by_percentage_share,
			ciaCountries.Territories[i].Budget,
			ciaCountries.Territories[i].Taxes_and_other_revenues,
			ciaCountries.Territories[i].Budget_surplus__or_deficit_,
			ciaCountries.Territories[i].Public_debt,
			ciaCountries.Territories[i].Fiscal_year,
			ciaCountries.Territories[i].Current_account_balance,
			ciaCountries.Territories[i].Exports,
			ciaCountries.Territories[i].Exports_partners,
			ciaCountries.Territories[i].Exports_commodities,
			ciaCountries.Territories[i].Imports,
			ciaCountries.Territories[i].Imports_partners,
			ciaCountries.Territories[i].Imports_commodities,
			ciaCountries.Territories[i].Reserves_of_foreign_exchange_and_gold,
			ciaCountries.Territories[i].Debt_external,
			ciaCountries.Territories[i].Exchange_rates,
			ciaCountries.Territories[i].Electricity_access,
			ciaCountries.Territories[i].Electricity_production,
			ciaCountries.Territories[i].Electricity_consumption,
			ciaCountries.Territories[i].Electricity_exports,
			ciaCountries.Territories[i].Electricity_imports,
			ciaCountries.Territories[i].Electricity_installed_generating_capacity,
			ciaCountries.Territories[i].Electricity_from_fossil_fuels,
			ciaCountries.Territories[i].Electricity_from_nuclear_fuels,
			ciaCountries.Territories[i].Electricity_from_hydroelectric_plants,
			ciaCountries.Territories[i].Electricity_from_other_renewable_sources,
			ciaCountries.Territories[i].Crude_oil_production,
			ciaCountries.Territories[i].Crude_oil_exports,
			ciaCountries.Territories[i].Crude_oil_imports,
			ciaCountries.Territories[i].Crude_oil_proved_reserves,
			ciaCountries.Territories[i].Refined_petroleum_products_production,
			ciaCountries.Territories[i].Refined_petroleum_products_consumption,
			ciaCountries.Territories[i].Refined_petroleum_products_exports,
			ciaCountries.Territories[i].Refined_petroleum_products_imports,
			ciaCountries.Territories[i].Natural_gas_production,
			ciaCountries.Territories[i].Natural_gas_consumption,
			ciaCountries.Territories[i].Natural_gas_exports,
			ciaCountries.Territories[i].Natural_gas_imports,
			ciaCountries.Territories[i].Natural_gas_proved_reserves,
			ciaCountries.Territories[i].Carbon_dioxide_emissions_from_consumption_of_energy,
			ciaCountries.Territories[i].Telephones_fixed_lines,
			ciaCountries.Territories[i].Telephones_mobile_cellular,
			ciaCountries.Territories[i].Telecommunication_systems,
			ciaCountries.Territories[i].Broadcast_media,
			ciaCountries.Territories[i].Internet_country_code,
			ciaCountries.Territories[i].Internet_users,
			ciaCountries.Territories[i].Broadband_fixed_subscriptions,
			ciaCountries.Territories[i].National_air_transport_system,
			ciaCountries.Territories[i].Civil_aircraft_registration_country_code_prefix,
			ciaCountries.Territories[i].Airports,
			ciaCountries.Territories[i].Airports_with_paved_runways,
			ciaCountries.Territories[i].Airports_with_unpaved_runways,
			ciaCountries.Territories[i].Pipelines,
			ciaCountries.Territories[i].Railways,
			ciaCountries.Territories[i].Roadways,
			ciaCountries.Territories[i].Waterways,
			ciaCountries.Territories[i].Ports_and_terminals,
			ciaCountries.Territories[i].Military_and_security_forces,
			ciaCountries.Territories[i].Military_expenditures,
			ciaCountries.Territories[i].Military_and_security_service_personnel_strengths,
			ciaCountries.Territories[i].Military_equipment_inventories_and_acquisitions,
			ciaCountries.Territories[i].Military_service_age_and_obligation,
			ciaCountries.Territories[i].Military_note,
			ciaCountries.Territories[i].Disputes_international,
			ciaCountries.Territories[i].Refugees_and_internally_displaced_persons,
			ciaCountries.Territories[i].Trafficking_in_persons,
			ciaCountries.Territories[i].Illicit_drugs,
			ciaCountries.Territories[i].Gallery,
			ciaCountries.Territories[i].Flag,
			ciaCountries.Territories[i].Map,
		)
		CheckError(e)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
