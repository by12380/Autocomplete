# Autocomplete

Autocomplete is a highly performant, scalable, and available service in the kubernetes environment to provide an application with type-ahead search word suggestions similar to Google search.

---

## Intro

**Project in progress (2020 - Present)**

This repo contains additional components that will showcase a full system using Autocomplete.

However, the ultimate goal of this project is to create an autocomplete service that will be portable and reusable with other kubernetes projects.

The data used for Autocomplete search words is based upon the frequency of past search queries collected from the web component.

## Tech Stacks

Go, Kubernetes, [Helm](https://helm.sh/) (K8 package manager), MongoDB (Log DB), [Minio](https://min.io/) (S3-compatible storage), [Argo](https://argoproj.github.io/) (Workflow manager), [Gin](https://github.com/gin-gonic/gin) (Web framework)

## Architecture

<div align="center">
<img src="https://github.com/by12380/Autocomplete/blob/master/docs/images/autocomplete-architecture.svg" width="900px">
</div>

## Dev Setup

1. [Install Go](https://golang.org/doc/install).

2. Install minikube and kubectl.
    - Follow the instructions in https://minikube.sigs.k8s.io/docs/start/ to install both.

3. [Install Helm](https://helm.sh/docs/intro/install/).

4. Clone this repository and navgiate to the project's root directory.

5. Run `go get` to install go module dependencies.

6. Run `minikube start` to start your kubernetes environment.

7. Run `helm install default ./chart` to deploy project to kubernetes.

8. (Optional) Run `minikube dashboard` to open Kubernetes dashboard and check for deployment status.


## Q&A

<details>
  <summary>How did you implement the Autocomplete service?</summary>
  
  ### Q: How did you implement the Autocomplete service?

  Trie was used as the data structure behind the Autocomplete service.
  
  ---
</details>

<details>
  <summary>How do you ensure the efficiency of the Autocomplete service?</summary>

  ### Q: How do you ensure the efficiency of the Autocomplete service?
  
  #### Answer:
  Since searching for all words matching a prefix in a trie has a time complexity of O(n), n being the number of nodes in the trie, the performace will suffer as the size of the trie grows.
  
  To ensure the efficiency of search, we modified the trie to store top K results at each node for its corresponding prefix.
  
  This will increase the space complexity to O(nk), where k is the number of top results we store.
  
  This will reduce the time complexity for searching words for a given prefix to O(1), and total time complexity for search operation would be reduced to O(l), where l is the length of the prefix (input keyword).
  
  A sacrifice of increased space for better time complexiity is a worth it tradeoff.
  
  ---
</div>
</details>

<details>
  <summary>How do you ensure the service is scalable to handle high throughput?</summary>

  ### Q: How do you ensure the service is scalable to handle high throughput?

  #### Answer:
  Since the Autocomplete service is read only, we can easily create replicas of the service to handle more request load.
  
  We can utilize the autoscaling feature that is supported by kubernetes natively.
  
  ---
</details>


<details>
  <summary>How do you ensure the service is scalable to handle a growing list of suggested search words?</summary>

  ### Q: How do you ensure the service is scalable to handle a growing list of suggested search words?

  #### Answer:
  As the size of the trie grows (growing list of suggested search words in our bank), it will eventually hit the memory limit for each pod instances within the service.
  
  To avoid holding all suggested search words in one app instance, we can split the search word data by the first letter alphabetically into groups, ex ([A-I], [J-R], [S-Z]).
  
  Thankfully, with the help of Helm templates, we can easily and dynamically create kubernetes resource by updating the configuration files used by Helm.
  
  ---
</details>

<details>
  <summary>How do you route search request to the correct shard of service?</summary>

  ### Q: How do you route search request to the correct shard of service?

  #### Answer:
  We placed a reverse proxy behind the autocomplete endpoints in our api gateway service. The reverse proxy will determine which shard of the autocomplete service to relay the request to by looking at which alphabet range the the first letter of the user's request query falls into. Since the api gateway is auto-scalable, it will not be a bottleneck to our autocomplete service.
  
  ---
</details>

<details>
  <summary>How is the bank of search words generated?</summary>

  ### Q: How is the bank of search words generated?

  #### Answer:
  When users make search requests (not request for autocomplete suggestions) to the web component, the requests are continously loggeed to the log DB. A cron job will then trigger our "update autocompelete service" workflow. This workflow is defined and made possible to run in a kubernetes environment by Argo Workfkow, and we defined the workflow as such:

  1. Trigger aggregator service:
    The aggregator service will read from our log DB and count the frequency of past search queries. This frequency data will then be uploaded as a file to our blob stoarge once the computation is complete. A successful operation of this step will trigger the next step.
  2. Trigger a re-deploy on our autocomplete service.
    When the autocomplete service is redeployed, the initiation of the servicee will first pull the frequency data from the blob storage and use it as the basis for constructing our trie.
  
  We used frequency of past search quries as both the dictionary of words used to populate our trie as well as the ranking for our top-K-result feature. However, the aggregator logic can be define in anyway we needed.

  ---
</details>

How do you ensure high availability for your service?

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

- Geographical-based search word suggestions
- Client-server connection with web-socket
- Multilingual support (UTF-8 chars)
