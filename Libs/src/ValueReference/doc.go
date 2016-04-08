/*
The ValueReference package implements functionality to operate with referece as with referent.

The typical use is to instantiate ValueReference(r) as wrapper for &v, and operate Get / Assign ReferentValue as you operate on v.
It is quite helpfull when you have to operate on some variable(or structure field) by referece as interface{}, but do not want to duplicate dereferencing and type switch/assertion code.

For.Ex. if you have to feed model fields to database/sql Scan() method, instead of making a new array of interface{}, scanning and copying values back into fields,
you could just make an array of ValueReference(r) containing referencies to the field, and let Scan directly into model fields by implemented Scanner assigning referent values.


*/
package ValueReference
