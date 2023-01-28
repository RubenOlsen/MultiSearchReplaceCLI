# Multi Search and Replace CLI

This is a small utility that is used to search a file for a variable number of (regular expression)
pattern and their corresponding replacements.

It is well suited when you need to have the following tasks performed:

* Convert a live (production) message (file) into a template.
* Anonymize, or masking, data in a file before sending it off, or reuse it.
* Convert a live (production) message (file) and change the content with random data.

## Usage

`$ msr -p fileWithPatterns.txt -i inputFileName.xml -o outputFileName.xml`

The above example will read a set of *search(es)* and the corresponding *replace(es)*, apply those on the 
_inputFileName.xml_ and write the output to the _outputFileName.xml_. 

## File containing search and replace pattern
The file containing the search and replace patterns will have a search and corresponding replace pattern on
each line divided by one or more tab-characters.

### Simple example with text replacing text
Given the following content in the patterns file
```
quick   slow
brown   blue
```
a file which contains
```The Quick Brown Fox Jumped Over The Fence``` will be changes into ```The slow blue Fox Jumped Over The Fence```

### Regular expression example with text replacements
Given the following content in the patterns file
```
<firstName>\w+</firstName>    <firstName>Daniel</firstName>
```
will change all occurrences of a string within the _firstName_ tag with _Daniel_.

### Variable data replacement using replacement functions
Given the following content in the patterns file
```
<accountNumber>\d+</accountNumber>    <accountNumber>#NUMBER(11)#</accountNumber>
```
will change all occurrences of _98780612345_ inside the _accountNumber_ tag with a random number containing
11 numerals.

Given the following content in the patterns file
```
<customerId>\d+</customerId>    <customerId>#COUNTER#</customerId>
```
will change all occurrences of any number  inside the _customerId_ tag with a sequence number starting 
with 1. This means that if there are 10 <customerId>...</customerId> tags in the input file, the output
file will then have <customerId>1</customerId> up to <customerId>10</customerId> instead of the existing
numbers.



## Supported replacement functions



