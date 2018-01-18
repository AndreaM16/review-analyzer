from vaderSentiment.vaderSentiment import SentimentIntensityAnalyzer
import sys

analyzer = SentimentIntensityAnalyzer()

sentence = sys.argv[1]

vs = analyzer.polarity_scores(sentence)
print(vs['compound'])