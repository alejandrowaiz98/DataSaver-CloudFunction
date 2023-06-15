Hi! If you are viewing the repos in order this should be the second one, and because the first one was an example of a CRUD api, now i will create this CF in order to use it across all future personal or looking-for-job related projects to avoid the recreation of a DB instance and all the logic behind in order to save data that i will need. Also works as a testimony that i know how to create and work with Cloud Functions, right? Thats double prize. 

So... basically will be:

ExternalApp <---> DataSaver

Where ExternalApp will use the following structure to communicate (via JSON):

    PubsubMessage -> 
    Attributes: Origen, Date, MessageID, DestinyCollection,
    Data: Message data in JSON format

And the DataSaver will respond with:

- In case of err: return err 
- In case of success: return nil