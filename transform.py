import pandas as pd
import random
# Read the CSV file
df = pd.read_csv('train.csv')

# Get first row (excluding label in first column)
first_image = df.iloc[random.randint(0,42002), 1:].values.reshape(28, 28)

# Display the image
import matplotlib.pyplot as plt
plt.imshow(first_image, cmap='gray')
plt.axis('off')
plt.show()