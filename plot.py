import matplotlib.pyplot as plt
from matplotlib import cm
from matplotlib.ticker import LinearLocator
import numpy as np
import scipy as sp


fig, ax = plt.subplots(subplot_kw={"projection": "3d"})

matrix = []
with open("matrix.txt", "r") as f:
    lines = f.readlines()
    for line in lines:
        matrix.append([int(i) for i in line[:-1].split(", ") if i])


H = len(matrix)
W = len(matrix[0])

def conv2d(a, f):
    s = f.shape + tuple(np.subtract(a.shape, f.shape) + 1)
    strd = np.lib.stride_tricks.as_strided
    subM = strd(a, shape = s, strides = a.strides * 2)
    return np.einsum('ij,ijkl->kl', f, subM)

def plot3d():
    x = np.linspace(-10, 10, W)
    y = np.linspace(-10, 10, H)
                    
    X, Y = np.meshgrid(x, y)

    Z = np.array(matrix)
    Z = Z % 2
    # Z = 1 / (Z + 1)
    # Z = np.log(Z)
    # Z = Z / np.max(Z) * 2
    # Z = np.exp(-np.power(Z - np.mean(Z), 2) / (2 * np.var(Z))) / np.std(Z)й

    # Z = sp.ndimage.filters.gaussian_filter(Z, [0.01,0.01], mode='mirror')
    for _ in range(20):
        Z = conv2d(Z, np.array([[0.5,0.5,0.5], [0.5, 1, 0.5], [0.5,0.5,0.5]]))
        Z = np.pad(Z, pad_width=1, mode='constant', constant_values=0)

    # Plot the surface.
    surf = ax.plot_surface(X, Y, Z, cmap=cm.coolwarm,
                        linewidth=0, antialiased=False)

    # Customize the z axis.
    # ax.set_zlim(-5.01, 5.01)
    # ax.zaxis.set_major_locator(LinearLocator(10))
    # # A StrMethodFormatter is used automatically
    # ax.zaxis.set_major_formatter('{x:.02f}')

    # Add a color bar which maps values to colors.
    fig.colorbar(surf, shrink=0.5, aspect=5)

    plt.show()


def plotWorm():
    Z = np.array(matrix)
    # Z = 1 / (Z + 1)
    Z = np.log(Z)
    Z = Z / np.max(Z) * 2

    Z = sp.ndimage.filters.gaussian_filter(Z, [3,2], mode='constant')

    # Z[Z < 0.6 and Z > 0.4] = 0
    for i in range(H):
        for j in range(W):
            if Z[i,j] > 0.9 or Z[i,j] < 0.0:
                Z[i,j] = 0
    plt.imshow(Z)
    plt.show()

plot3d()