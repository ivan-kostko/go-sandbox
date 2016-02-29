
   -- DROP TABLE dbo.TypesTest

IF OBJECT_ID('dbo.TypesTest') IS NULL
   CREATE TABLE dbo.TypesTest (
      Col_INT             INT              ,
      Col_BIGINT          BIGINT           ,
      Col_DECIMAL         DECIMAL(25,7)    ,
      Col_BIT             BIT              ,
      Col_NUMERIC         NUMERIC(29,7)          ,
      Col_FLOAT           FLOAT            ,
      Col_DATE            DATE             ,
      Col_TIME            TIME             ,
      Col_DATETIME        DATETIME         ,
      Col_DATETIME2       DATETIME2        ,
      Col_SMALLDATETIME   SMALLDATETIME    ,
      Col_DATETIMEOFFSET  DATETIMEOFFSET   ,
      Col_NVARCHAR        NVARCHAR(4000)    ,
      --Col_XML             XML              ,
      Col_VARCHAR         VARCHAR(MAX)     ,
      Col_VARBINARY       VARBINARY(MAX)        
   )
GO

INSERT INTO dbo.TypesTest (
     Col_INT            
   , Col_BIGINT         
   , Col_DECIMAL        
   , Col_BIT            
   , Col_NUMERIC        
   , Col_FLOAT          
   , Col_DATE           
   , Col_TIME           
   , Col_DATETIME       
   , Col_DATETIME2      
   , Col_SMALLDATETIME  
   , Col_DATETIMEOFFSET 
   , Col_NVARCHAR       
   --, Col_XML          
   , Col_VARCHAR        
   , Col_VARBINARY      
)
SELECT 
   2147483646 AS [Col_INT]
   , 9223372036854775806 AS [Col_BIGINT]
   , 12345.1234567 AS [Col_DECIMAL]
   , 1 AS [Col_BIT]
   , 12345.1234567 AS [Col_NUMERIC]
   , 123456789.123456789 AS [Col_FLOAT]
   , N'2016-01-31' AS [Col_DATE]
   , N'09:05:33.1100000' AS [Col_TIME]
   , N'2016-01-31 09:05:33.110' AS [Col_DATETIME]
   , N'2016-01-31 09:05:33.1100000' AS [Col_DATETIME2]
   , N'2016-01-31 09:05:00' AS [Col_SMALLDATETIME]
   , N'2016-01-31 09:05:33.1100000 +00:00' AS [Col_DATETIMEOFFSET]
   , N'driver=SQL Server;server=.\MSSQL_2014_DEV;uid=godev;pwd=godev;database=Test' AS [Col_NVARCHAR]
   , 'driver=SQL Server;server=.\MSSQL_2014_DEV;uid=godev;pwd=godev;database=Test' AS [Col_VARCHAR]
   , 0x6472697665723D53514C205365727665723B7365727665723D2E5C4D5353514C5F323031345F4445563B7569643D676F6465763B7077643D676F6465763B64617461626173653D54657374 AS [Col_VARBINARY] 
   
SELECT * FROM dbo.TypesTest   
