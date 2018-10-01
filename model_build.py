
# coding: utf-8

# In[17]:

import pandas as pd 
from sklearn.model_selection import train_test_split
from sklearn.neural_network import MLPClassifier
from sklearn.metrics import accuracy_score


# In[77]:

data = pd.read_csv("gcskewdata_100_5.csv", sep=";", header=None)

data.shape


# In[82]:

y = data.iloc[:,-1]
X = data.iloc[:, :-1]

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.20)

clf = MLPClassifier()

clf.fit(X_train, y_train)

accuracy_score(y_test, clf.predict(X_test))


# In[72]:



