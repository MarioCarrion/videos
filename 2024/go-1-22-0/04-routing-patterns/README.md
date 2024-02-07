# Enhanced routing patterns

Excerpt from official release notes:

> The patterns used by net/http.ServeMux now accept methods and wildcards: For example, the router accepts a pattern
> like GET /task/{id}/, which matches only GET requests and captures the value of the {id} segment in a map that can be
> accessed through Request values.
