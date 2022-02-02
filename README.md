# Produce Data

This repo will generate fake data for the [SurveyFunky](https://github.com/arelangi/SurveyFunky) mock app.

- [GoFakeIt](https://github.com/brianvoe/gofakeit) library will be used to generate the data


We will start working on the User object

- [x] Generate data to DB
- [x] Publish data to Kafka
	- [x] Register schema in schema registry
	- [x] Write to Kafka
- [x] Ensure data lands in Hudi
- [x] Ensure data is queryable by ad-hoc analysis tools

--------------

02/02/2022

- [ ] Refactor User.go and abstract away the push to Kafka
- [x] Setup a proper partition path to be able to query data (Test non-date partitions)


