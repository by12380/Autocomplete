# Autocomplete

Autocomplete is a highly performant, scalable, and available service in the kubernetes environment to provide an application with type-ahead search word suggestions similar to Google search.

---

## Intro

This repo contains additional components that will showcase a full system using Autocomplete.

However, the ultimate goal of this project is to create an autocomplete service that will be portable and reusable with other kubernetes projects.

The data used for Autocomplete search words is based upon the frequency of past search queries collected from the web component.

## Tech Stacks

Go, Kubernetes, [Helm](https://helm.sh/) (K8 package manager), MongoDB (Log DB), [Minio](https://min.io/) (S3-compatible storage), [Argo](https://argoproj.github.io/) (Workflow manager), [Gin](https://github.com/gin-gonic/gin) (Web framework)

## Architecture

<div align="center">
<img src="https://github.com/by12380/Autocomplete/blob/master/docs/images/autocomplete-architecture.svg" width="900px">
</div>

## Q&A

How did you implement the Autocomplete service?

How do you ensure the efficiency of the Autocomplete service?

How do you ensure the service is scalable to handle high throughput?

How do you ensure the service is scalable to handle a growing list of suggested search words?

How did you route search request to the correct shard of service?

How do you ensure high availability for your service?

How is the bank of search words generated?

How are the search words in the service updated in real time?

How did you cooordinate sequence of actions to happen in Kubernetes?

How do you prevent downtime when the service is being updated?

Why did you use Helm? What are the benefits over using kubernetes alone?

How do you re-shard the service?

Why did you use mongoDB as a log database?

How did you ensure logging does not slow down the speed of your service?

What is Minio?

Why did you not use exisitng cloud services such as AWS S3 and Mongo Atlas?

## Next Steps