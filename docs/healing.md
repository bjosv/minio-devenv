# Healing tests

## Commands
kubectl exec -it pod/minio-0 -- rm -rf /export/bucket-1
kubectl exec -it pod/minio-0 -- rm -rf /export/bucket-1 /export/bucket-2 /export/bucket-3 /export/bucket-4 /export/bucket-5 /export/bucket-6 /export/bucket-7 /export/bucket-8 /export/bucket-9 /export/bucket-10

kubectl delete pods minio-1 --grace-period=0 --force

#### Newer
mc admin info local
time mc admin heal --recursive local

#### Older
mc-2020-02-20 admin info local
time mc-2020-02-20 admin heal --recursive local

9.4 GiB Used, 10 Buckets, 16,000 Objects

## XPs

### RELEASE.2021-06-17T00-10-46Z
1,600/16,000 objects; 9 GiB in 1m25s
1,598/16,000 objects; 9 GiB in 1m23s
14,411/16,000 objects; 9 GiB in 7m9s
13,925/16,000 objects; 9 GiB in 7m26s


### RELEASE.2021-06-14T01-29-23Z but with removed fsync
1,599/16,000 objects; 9 GiB in 1m3s
15,993/16,000 objects; 9 GiB in 5m29s

### RELEASE.2021-06-14T01-29-23Z
mc
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/16,000 objects; 9 GiB in 3m7s
1,598/16,000 objects; 9 GiB in 1m8s
1,600/16,000 objects; 9 GiB in 1m10s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
14,342/16,000 objects; 9 GiB in 7m52s
15,306/16,000 objects; 9 GiB in 6m16s
15,374/16,000 objects; 9 GiB in 6m36s

# No-action heal
0/16,000 objects; 9 GiB in 2m38s
0/16,000 objects; 9 GiB in 2m49s
0/16,000 objects; 9 GiB in 33s
0/16,000 objects; 9 GiB in 33s

# Kill pod

# No persistence
```

### RELEASE.2021-02-14T04-01-33Z
mc
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 56s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
15,842/16,000 objects; 9 GiB in 9m31s
```

### RELEASE.2021-01-16T02-19-44Z
mc - 4drives
1,596/16,000 objects; 9 GiB in 2m55s
1,600/16,000 objects; 9 GiB in 2m54s
15,644/16,000 objects; 9 GiB in 11m16s

mc-2020-02-20


### RELEASE.2020-08-27T05-16-20Z
mc-2020-02-20
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/16,000 objects; 9 GiB in 3m34s   mc: checks all buckets
1,600/16,000 objects; 9 GiB in 7m44s   mc: checks all buckets
1,600/1,600 objects; 961 MiB in 1m14s  mc-2020-02-20

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 13m33s
15,997/16,000 objects; 9 GiB in 14m45s
15,997/16,000 objects; 9 GiB in 12m13s   mc-2020-02-20
```

### RELEASE.2020-07-22T00-26-33Z
mc-2020-02-20 admin info local
time mc-2020-02-20 admin heal --recursive local
```
9.4 GiB Used, 10 Buckets, 16,000 Objects

# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 1m2s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 10m13s
```

RELEASE.2020-07-11T06-07-16Z - Major Feature Release

RELEASE.2020-07-11T21-14-23Z
RELEASE.2020-07-11T06-07-16Z

#### got worse..

### RELEASE.2020-07-02T00-15-09Z
mc-2020-02-20 admin info local
time mc-2020-02-20 admin heal --recursive local
```
7.5 GiB Used, 10 Buckets, 13,471 Objects

# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 50s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 8m12s
```

### Test of 12e691a7a
Commit: add additional fdatasync before close() on writes (#9947)
https://github.com/minio/minio/pull/9947
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 45s
1,600/1,600 objects; 961 MiB in 47s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 8m19s
16,000/16,000 objects; 9 GiB in 8m7s
```

#### got worse..

### Test of 9b1876c56
Commit before: add additional fdatasync before close() on writes (#9947)
mc-2020-02-20 admin info local
time mc-2020-02-20 admin heal --recursive local
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 40s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 5m45s
```


### RELEASE.2020-06-22T03-12-50Z
mc-2020-02-20 admin info local
time mc-2020-02-20 admin heal --recursive local
```
6.6 GiB Used, 10 Buckets, 11,772 Objects

# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 36s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 5m37s
```

### RELEASE.2020-05-28T23-29-21Z
mc-2020-02-20
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 34s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 5m36s
```

### RELEASE.2020-02-20T22-51-23Z
mc-2020-02-20
```
mc-2020-02-20 admin info local
time mc-2020-02-20 admin heal --recursive local

# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 37s
1,600/1,600 objects; 961 MiB in 36s
1,600/1,600 objects; 961 MiB in 35s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 5m53s
16,000/16,000 objects; 9 GiB in 5m42s

# No-action heal
0/0 objects; 0 B in 1s

# Kill pod
mc-2020-02-20: <ERROR> Failed to start heal sequence. Heal is already running on the given path (use force-start option to stop and start afresh). The heal was started by IP 127.0.0.1 at Fri, 18 Jun 2021 11:44:03 GMT, token is cfec578e-6526-4c84-b8d4-c0b7c857abd5
```

https://gist.github.com/harshavardhana/a8d1669345fd4c5d388447ab69520fcc


# Issues
2021-03: Healing a disk is taking a very long time
https://github.com/minio/minio/issues/11696
