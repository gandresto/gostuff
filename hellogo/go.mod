module example.com/gandresto/gostuff/hellogo

go 1.21.5

// Typically we dont see this replace directive in production
replace example.com/gandresto/gostuff/mystrings v0.0.0 => ../mystrings

require example.com/gandresto/gostuff/mystrings v0.0.0