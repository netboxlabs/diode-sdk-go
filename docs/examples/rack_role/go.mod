module rack_role

go 1.25

require github.com/netboxlabs/diode-sdk-go

// Use local SDK for development and validation
// When copying this example, remove the replace directive and update the require version above
replace github.com/netboxlabs/diode-sdk-go => ../../..
