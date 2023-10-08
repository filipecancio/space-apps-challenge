# Desafio Space Challenge
[![badge](https://img.shields.io/badge/prototipo-figma-red)](https://www.figma.com/file/VL89dWkK1FKHlTRRFwUwiR/SpaceApps2023?type=design&node-id=0%3A1&mode=design&t=xzj6eWJx1QdqRAWd-1)
[![badge](https://img.shields.io/badge/fluxograma-figma-blue)](https://www.figma.com/file/qwLEcIzg0By7YxGsNKIB6T/Fluxos?type=whiteboard&node-id=0%3A1&t=iPoKMpisyWOd5RBk-1)
[![badge](https://img.shields.io/badge/app-prototype-red)](https://www.figma.com/proto/VL89dWkK1FKHlTRRFwUwiR/Prot%C3%B3tipo?page-id=0%3A1&type=design&node-id=25-1042&viewport=-2184%2C-33%2C0.44&t=vq7m7UTxUQgQOjnM-1&scaling=min-zoom&starting-point-node-id=25%3A1042&mode=design)
[![badge](https://img.shields.io/badge/chatbot-prototype-red)](https://www.figma.com/proto/VL89dWkK1FKHlTRRFwUwiR/Prot%C3%B3tipo?page-id=0%3A1&type=design&node-id=99-1704&viewport=-2184%2C-33%2C0.44&t=vq7m7UTxUQgQOjnM-1&scaling=min-zoom&starting-point-node-id=99%3A1704&show-proto-sidebar=1&mode=design)

Desafio Escolhido: [Gerenciamento de incêndios: aumentando as oportunidades de gerenciamento de incêndios com base na comunidade](https://github.com/filipecancio/space-apps-challenge/issues/1)

## High-Level Summary
We developed an API that integrates the georeferencing of Odin Fire to aggregate the community in the generation of fire alert data. The community accesses our system through a chatbot that is easy and practical to use to report possible fires. With it, we can send the user's current location, as well as sending images of the location and describing the fire scenario. However, using a chatbot system today, whether via WhatsApp or Telegram, requires access to the local internet, making it impossible to focus on this routine in rural areas, far from urban centers.  
To this end, an easy and practical application was developed that would use the same API, with which it is possible to issue a local fire alert with just two clicks, in which case it would be easier to use for users without an understanding of technology, and for regions without internet, the app has internal data storage, providing data such as the current location, and from an internet connection or telephone signal, its data is sent by an online request or by an SMS.
So the community will have records of fires provided by geolocation in an accessible way and will also be able to contribute to the provision of the same, solving the problem in question consists of addressing the monitoring of fires and natural resources through the innovative use of technology and publicly available data, in order to make it possible for local communities to report and monitor fires and/or improve the current distribution of data.
The provision of a public API is important because it allows any citizen to easily access the data provided both by the contribution of others and by data already identified via satellite.
In addition, the application developed would allow a local emergency call to be made when necessary, avoiding the spread of smaller fires.

## Project Details
Our project is made up of three interconnected features: The API that correlates the georeferencing data provided by Nasa through ODIN, a chatbot that makes a submission request to report a fire, and a complete application that with the data received from the API will show reported fires on a map and will also make it possible to submit a report, this time with options for images, observations and more details of the fire, with some tabs of the APP also being geared towards awareness and emergency calls. 

Its benefit is to add more data to NASA's existing tools for identifying fires, as well as democratizing access within communities, bringing accessibility to people with access to cell phones.

It is hoped to achieve a powerful tool that can be developed to support large-scale data and receive attention from public bodies at national and regional levels to motivate communities to use it, with a view to possible implementation of a general data management system.

Figma was used to design the application screens, Git and GitHub for general project management, GO and Gin Web Framework for API development, Kotlin for mobile application development and TyperScript for ChatBot development.

## Use of Artificial Intelligence

We used the Open AI API in conjunction with the Telegram chatbot API to generate the chat assistant. We also used AI to help translate texts quickly, with DeepL.
