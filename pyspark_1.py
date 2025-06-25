# Databricks notebook source
# No need to import SparkSession or create the spark object in Databricks
from pyspark.sql import SparkSession
spark=SparkSession.builder.appName('nigesh').getOrCreate()
data=[(1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
       (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano'),
      (1,'nigesh'),(2,'praveen'),(3,'yogesh'),(4,'mano')
      ]
schema=['id','name']
df=spark.createDataFrame(data,schema)
display(df)

# COMMAND ----------

from pyspark.sql import SparkSession
spark=SparkSession.builder.appName('nigesh').getOrCreate()
df=spark.read.option('header',True).option('inferschems',True).csv('/Volumes/workspace/default/nigesh/car_prices.csv').display()

# COMMAND ----------

df = spark.read.csv("/Volumes/workspace/default/nigesh/car_prices.csv", header=True, inferSchema=True)
df.select("make").display()

# COMMAND ----------

from pyspark.sql.functions import col
df.select(col("make")).display()

# COMMAND ----------

columns=[col for col in df.columns]
df.select(*columns).display()

# COMMAND ----------

df.select(df.columns).display()

# COMMAND ----------

df.select(df.columns[0:5]).display()

# COMMAND ----------

df.select([df["make"],df["model"]]).display()

# COMMAND ----------


df.select(df['make'],df['model']).display()

# COMMAND ----------

df.select(col('make')).display()

# COMMAND ----------

columns=[col for col in df.columns]
df.select(*columns).display()

# COMMAND ----------

df.select(df.columns).display()

# COMMAND ----------

df.select(df.columns[0:4]).display()

# COMMAND ----------

df.select([df["make"],df["model"]]).display()

# COMMAND ----------

from pyspark.sql.functions import col,lit,concat_ws
df2=df.withColumn('total price',col('sellingprice')+lit(1000))
df2.select(df2.columns).display()

# COMMAND ----------

from pyspark.sql.functions import col,lit,concat_ws
from pyspark.sql import SparkSession
spark=SparkSession.builder.appName('nigesh').getOrCreate()
df=spark.read.option('header',True).option('inferschems',True).csv('/Volumes/workspace/default/nigesh/car_prices.csv')
df2=df.withColumn('total price',col('sellingprice')+lit(1000))
df2.select(df2.columns)
df2=df.withColumn('total price',col('sellingprice')+lit(1000))
df2=df.withColumn('total price',col('sellingprice')-lit(100))
df3=df2.withColumn('model',concat_ws(' ',col('make'),col('model')))
df4=df3.drop('make')
df5=df.collect()
print(df5)

# COMMAND ----------

from pyspark.sql import SparkSession
spark=SparkSession.builder.appName('nigesh').getOrCreate()
df=spark.read.option('header',True).option('inferschema',True).csv('/Volumes/workspace/default/nigesh/car_prices.csv')

# COMMAND ----------

df5=df.collect()
df6=df5[0][5]
print(df6)