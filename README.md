# Autocomplete

Autocomplete is a highly performant, scalable, and available service in the kubenetes environment to provide an application with type-ahead search word suggestions similar to Google search.

---

## Intro

This repo contains additional components that will showcase a full system using Autocomplete.

However, the ultimate goal of this project is to create an autocomplete service that will be portable and reusable with other kubernetes projects.

The data used for Autocomplete search words is based upon the frequency of past search queries collected from the web component.

## Tech Stacks

Go, Kubenetes, [Helm](https://helm.sh/) (K8 package manager), MongoDB (Log DB), [Minio](https://min.io/) (S3-compatible storage), [Argo](https://argoproj.github.io/) (Workflow manager), [Gin](https://github.com/gin-gonic/gin) (Web framework)

## Architecture

<div align="center">
<img src="https://github.com/by12380/Autocomplete/blob/master/docs/images/autocomplete-architecture.svg" width="900px">
</div>
