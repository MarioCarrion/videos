# Anti-corruption Layer Pattern

This example represents a system that returns temperature for different countries, in our domain we use Celsius.

We are re-using two third party APIs represented as local packages in:

* `fahrenheit`
* `kelvin`

Because those upstream packages represent third party APIs (or maybe legacy systems):

* They may define their own domain, which in this case is represent as different temperature types, and
* They may be implemented in a way that is not idiomatic, follows best practices or does not match our guidelines,

So in order to not _corrupt_ our system we implement a layer to handle converting the data we receive from them and the data we need to send them, to keep those packages isolated from our implementation.
