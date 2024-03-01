import numpy as np
import matplotlib.pyplot as plt
from gensim.models import Word2Vec
from gensim.utils import simple_preprocess

documents = [
    'the quick brown fox jumps over the lazy dog',
    'this is my example sentance'
]

processed_docs = [simple_preprocess(doc) for doc in documents]
model = Word2Vec(sentences=processed_docs, vector_size=100, window=5, min_count=1, workers=4)
model.save("word2vec.model")

def encode_document(doc):
    return np.mean([model.wv[word] for word in doc if word in model.wv], axis=0)

encoded_doc = encode_document(processed_docs[0])

# Example 100-dimensional vector (replace with your actual vector)
vector = np.array(encoded_doc)  # Using random data for illustration
vector_length = vector.shape[0]
side_length = int(np.ceil(np.sqrt(vector_length)))

square_matrix = np.zeros((side_length * side_length))
square_matrix[:vector_length] = vector

square_matrix = square_matrix.reshape((side_length, side_length))

# Reshape the vector into a 10x10 grid

# Create an image
plt.imshow(square_matrix, cmap='hot', interpolation='nearest')
plt.axis('off')  # Hide axes for a cleaner look
plt.show()

