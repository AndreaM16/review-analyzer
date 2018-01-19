# review-analyzer
Golang + Python Sentiment Analsys Utility

## Project's Purpose
Given a set of Amazon Reviews, not in consecuitve days, and some days may have multiple reviews:

 - Performs a Polarized Sentiment Analysis on each review's body
 - Flattens the reviews per each day, so, we will have only one review per day and If there are more entries per each days, stars and sentiment is the average of them
 - Fills the missing days between each entry with an average of the entries
 - Posts the results to api
