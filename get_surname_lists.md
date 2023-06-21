#US Surname lists
#zipped Excel and CSV (comma separated) files of the complete list of 162,253 names
curl -O https://www2.census.gov/topics/genealogy/2010surnames/names.zip
#Extract just what we want and format it
unzip -p names.zip Names_2010Census.csv | sed '1d;$d' | awk -F',' '{print tolower($1)}' > surnames-us-ordered.txt



From 2010 census data:

https://www.census.gov/topics/population/genealogy/data/2010_surnames.html