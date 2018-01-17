from vaderSentiment.vaderSentiment import SentimentIntensityAnalyzer

analyzer = SentimentIntensityAnalyzer()

sentence = sys.argv[1]

vs = analyzer.polarity_scores(sentence)
print(vs['compound'])