# Exported variables can be replaced

Because exported variables can be replaced by anyone importing the package you should consider the following:

* When using packages, **do not** use the instance of exported variable directly, instead instantiate your own,
* When implementing packages:
    * If the variable is a **"basic type"** use a constant instead,
    * If you want to return a default implementation, consider creating a method `Default()` instead of a exported variable `Default`.

## Examples

* [v1](v1/): Uses typical exported variables, and
* [v2](v2/): Uses my recommendations.
