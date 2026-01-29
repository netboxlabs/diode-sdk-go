# Generating Diode Entity Documentation from Diode Code

Currently the existing Diode documentation on using Diode entities is here: https://github.com/netboxlabs/diode-sdk-python/blob/5dae6c3222fb99124c3e22bd312ebc4a537ccf9f/docs/entities.md

That was manually created and because of the volume of entities and the constant change we want to automate that.

My understanding is that the path to this is via the Diode data serialization code (in Python), that is then converted to Protobuf specifications, which is then generated docs

Note that the diode repo is in the same parent directory as this project, at /Users/pstuart/CODE/github.com/netboxlabs/diode -- in fact all diode related projects are in this parent directory for local search.

Let's start by identifying any code in /Users/pstuart/CODE/github.com/netboxlabs/diode-sdk-python that analyzes diode django marshalers for translating.

Create @ENTITIES_PLAN.md and update it with:

* your understanding of the task
* existing code that analyzes/translates the django marshalers
* plan how generate the diode-sdk-go documentation 