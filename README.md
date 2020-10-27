# Autocomplete

Autocomplete is a highly performant, scalable, and available service for providing an application with type-ahead search word suggestions similar to Google seearch.

---

An additional goal of this project is to create an autocomplete service that is portable and reusable, decoupled from the rest of the system.

However, to demonstrate it's use case, this repo contains the rest of the components that will represent a full system using Autocomplete.

The data used for Autocomplete search words is based upon the frequency of past search queries from a separate web component.

---

## Architecture

<div align="center">
<img src="https://github.com/by12380/Autocomplete/blob/master/docs/images/autocomplete-architecture.svg" width="600px">
</div>