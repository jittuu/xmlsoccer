/*
Package xmlsoccer provides the proxy client to call xmlsoccer webservice. (http://www.xmlsoccer.com/)

Input parameters

league
  The league parameter is a string and can either be the alphanumeric complete name, or the numeric ID.
  For example, "English Premier League" or "1"

team
  The team parameter is a string and can either be the alphanumeric complete name, or the numeric ID.
  For example, "Inverness CT" or "1"

season
  The season parameter is the last two digits of the beggining of the season-year appended by the last two digits of the following year.
  For example,
    "1213" for the season of 2012-2013
    "0809" for the season of 2008-2009
  The American "Major league" and the Swedish "Allsvenskan" is two examples of this,
  as they start in the beginning of the year and end in the end of the year.
  Despite of this, they too follow the same seasonDateString rules.
  So the 2013 season of the American "Major League" will need the input: "1314"
*/
package xmlsoccer
