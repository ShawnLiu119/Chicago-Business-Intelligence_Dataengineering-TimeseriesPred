# Data-Engineering-Chicago-Business-Intelligence-Report
Data-Engineering-Chicago-Business-Intelligence-Report


### Overview
COVID-19 pandemic outbreak has brought great challenge to both social economy and public health management of U.S. in past two years. In the process of monitoring and assessing the pandemic hit to local communities, big data has played an important role by enabling timely tracking of both economic and public health metrics such as transportation, unemployment, daily cases/ deaths and so on. The city of Chicago has laid out data infrastructure with 16 categories of datasets in place, this project aims to build data lake based on 3 main categories including Transportation, Buildings, and Health & Human Services, to enable capability of generating different intelligence reports for various purposes such as strategic planning, disease control, and infrastructure investment. 

### 1. Data Source and Collection
https://data.cityofchicago.org/. 

Please refer more details to the final report

### 2. Infrastructure and Technology Stack

2.1 Database Engine

PostgreSQL is a powerful, open-source object-relational database system that uses and extends the SQL language combined with many features that safely store and scale the most complicated data workloads. PostgreSQL 14 brings a variety of features that help developers and administrators deploy their data-backed applications. PostgreSQL continues to add innovations on complex data types, including more convenient access for JSON and support for noncontiguous ranges of data.

2.2 Micro Service Development

Python is used as programming language in this project to construct the microservices which will be tasked to extract data from Chicago Data Portal, interact with Postgres Engine and insert data into database. A few python packages are utilized as key components to execute functionalities including data extraction, ingestion, and transformation. Golang has also been reviewed as alternative language which provides multiple advantages over Python. 

2.3 Micro Service Deployment

Docker is an excellent tool for managing and deploying microservices. Each microservice can be further broken down into processes running in separate Docker containers, which can be specified with Dockerfiles and Docker Compose configuration files.
Kubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of containerized applications

2.4 Time Series Forecast

To develop time-series forecast capability of projecting taxi traffic patterns, TensorFlow LSTM algorithm from Google and Prophet from Facebook are under consideration.
TensorFlow/Keras LSTM, Long Short-Term Memory layer, is one of Recurrent neural networks (RNN) that is powerful for modelling sequence data such as time series. Schematically, a RNN layer uses a for loop to iterate over the timesteps of a sequence, while maintaining an internal state that encodes information about the timesteps it has seen so far. Prophet is an open source library released by Facebook. It delivers good performance, especially when working with time series that have strong seasonal effects fitting with non-linear trends. 

### 3. Architecture and Design

<img width="468" alt="image" src="https://user-images.githubusercontent.com/43327902/185725710-13df9a91-c3ed-41d8-8e4f-d5201e5786a2.png">

3.1 Data Collection and Collection

Data Ingestion and collection is designed as and fulfilled by the microservice built with Python API, containerized by Docker, and deployed via Kubernetes resources. The python app is built with data extraction and insertion functionalities that interact with data lake on Postgres Engine hosted on local machine. To implement the functionalities, several Python libraries are utilized and built into the app:
-	Pandas: a widely used for data processing and manipulation.
-	Sodapy: python client for the Socrata Open Data API to communicate with Chicago Data Portal (https://github.com/xmunoz/sodapy)
-	Psycopy2: one of popular PostgreSQL database adapter for python that is used to interact with data lake constructed on Postgres engine
-	Schedule: package for job scheduling (https://schedule.readthedocs.io/en/stable/)
Then, the app is put into a container via Docker and deployment via Kubernetes. To avoid technical issues with connection, a virtual environment has been initiated via minikube. The image built into container will be pushed to Docker Hub and deployed by Kubernetes. The microservice is scheduled on daily basis with automatic execution

3.2 Data Storage

The data storage layer is built upon a data lake on PostgreSQL engine. The fragmented data collected from Chicago Data Portal will be stored into “chiproject_postgres” database hosted on local machine (Port:5432). Within the data base, six tables are created as show below to store data that is extracted with selected fields from respective API regarding each data categories described in phase 1. 

3.3 Data Processing 
In this layer, the fragmented data that is stored in PostgreSQL database is queried with SQL statements and loaded into Jupyter notebook using Psycopy2 depending on the specific requirements. The data is then manipulated and processed through several foundational steps such as removing null value, mapping out zip code with community area number, mapping location coordinates (latitude, longitude) to zip code, etc. Cleaned data with lean relevant fields are expected to be generated ready for visualization and forecast in next Insight layer. 

3.4 Insights and visualization

In this layer, analytical algorithms such as TensorFlow LSTM algorithm from Google and Prophet from Facebook are brought in and used for time series forecast on traffic pattern in daily, weekly and monthly window. Also, some tabular-format output is generated to meet the requirements by City of Chicago, for example, the top 5 neighborhoods with highest permits but highest unemployment rate and poverty rate are identified and delivered as one of projects outputs. 


### 4. Development

4.1 Data lake construction

The data lake that hosts the raw data extracted from Chicago Data Portal was constructed with Postgres Database engine. The database and tables were defined through pgAdmin4 and hosted on local machine (port 5432) as shown in Figure 2. 

<img width="468" alt="image" src="https://user-images.githubusercontent.com/43327902/185725759-c0fbe47e-149d-4f37-88b9-5d4a0492595a.png">

4.2 Data fetching API
The API that is tasked to fetch data from Chicago Data Portal in a batch manner was developed with Python environment and containerized and deployed using Docker/Kubernetes. All the coding was done in Visual Studio Code, which integrates well with different programming language. As shown in Figure.3, the data fetching functionalities was programed in main.py, then it was containerized into Docker image (Figure.4) and further defined and run via yaml file (Figure.5). 

<img width="464" alt="image" src="https://user-images.githubusercontent.com/43327902/185725767-c746ecbc-af0a-4d88-988a-c6edf1ee9949.png">

<img width="466" alt="image" src="https://user-images.githubusercontent.com/43327902/185725770-a68c8dcb-723b-4eaa-9ec6-4da9e4937478.png">

<img width="470" alt="image" src="https://user-images.githubusercontent.com/43327902/185725774-87fd55e5-2eb9-416b-9727-ebb83f171668.png">

<img width="468" alt="image" src="https://user-images.githubusercontent.com/43327902/185725776-3746c6f0-846c-4ed8-accb-6189c17c3aa9.png">

4.3 Data processing and analysis
Data processing and analysis were conducted in Python environment. The functionalities that were developed and implemented in this section are purposed to meet the described requirements and enable City of Chicago to make strategic planning with analytical and forecasting capabilities. 
More details please refer to the report

### 5. Forecasting and Strategic Planning

As described in section 4 – requirement 4, the forecast of traffic pattern has been generated using Prophet model at daily, weekly, and monthly basis. Due to computing intensity driven by large data size, 2021 taxi trip date was extracted from date lake and used for demo. Also, stratified sampling method has been used to proportionally sample the taxi trip data by date. 

The figures below demonstrate the output of times series forecast by Prophet model. 

<img width="414" alt="image" src="https://user-images.githubusercontent.com/43327902/185725828-7f3e6bcb-4726-4f82-aed8-9b0a9efa9b4e.png">

Figure 12 Daily Traffic Forecast – Zip Code 60657

<img width="408" alt="image" src="https://user-images.githubusercontent.com/43327902/185725830-93e7705c-e1d0-4c33-bf49-5e061c9bc2bf.png">

Figure 13 Weekly Traffic Forecast – Zip Code 60657

<img width="422" alt="image" src="https://user-images.githubusercontent.com/43327902/185725832-12cd11ce-a584-49fb-adbf-9775a7883730.png">

Figure 14 Monthly Traffic Forecast – Zip Code 60657

<img width="408" alt="image" src="https://user-images.githubusercontent.com/43327902/185725834-852e8613-93c4-4417-9067-a67a7a01a013.png">

Figure 15 Daily Traffic Forecast – All Zip Codes

<img width="425" alt="image" src="https://user-images.githubusercontent.com/43327902/185725860-5e9cbc54-dfae-46f6-b5c2-44eb7ce808fc.png">

Figure 16 Weekly Traffic Forecast – All Zip Codes!

<img width="421" alt="image" src="https://user-images.githubusercontent.com/43327902/185725864-dbd5347f-5a9b-431c-8f96-f6139e476162.png">

Figure 17 Monthly Traffic Forecast – All Zip Codes









