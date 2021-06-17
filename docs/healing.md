# Healing tests

## Commands
### Remove files commands
kubectl exec -it pod/minio-0 -- rm -rf /export/bucket-1
kubectl exec -it pod/minio-0 -- rm -rf /export/bucket-1 /export/bucket-2 /export/bucket-3 /export/bucket-4 /export/bucket-5 /export/bucket-6 /export/bucket-7 /export/bucket-8 /export/bucket-9 /export/bucket-10

### Trigger heal using latest `mc` (v3 API)
mc admin info local
time mc admin heal --recursive local

### Trigger heal using older `mc` (v2 API)
mc-2020-02-20 admin info local
time mc-2020-02-20 admin heal --recursive local

## Experiments

### RELEASE.2021-06-17T00-10-46Z
Using mc
```
1,600/16,000 objects; 9 GiB in 1m25s
1,598/16,000 objects; 9 GiB in 1m23s
14,411/16,000 objects; 9 GiB in 7m9s
13,925/16,000 objects; 9 GiB in 7m26s
```

### RELEASE.2021-06-14T01-29-23Z but with removed fdatasync
Using mc
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,599/16,000 objects; 9 GiB in 1m3s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
15,993/16,000 objects; 9 GiB in 5m29s
```

### RELEASE.2021-06-14T01-29-23Z
Using mc
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,599/16,000 objects; 9 GiB in 1m13s
1,596/16,000 objects; 9 GiB in 1m12s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
15,497/16,000 objects; 9 GiB in 6m32

?? -> 8.2 GiB Used, 10 Buckets, 13,989 Objects
```

### RELEASE.2021-02-14T04-01-33Z
Using mc
```
# Single bucket (1 bucket, 800x10kb, 400x500kb, 400x2MB)
1,600/1,600 objects; 961 MiB in 56s

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
15,842/16,000 objects; 9 GiB in 9m31s
```

### RELEASE.2021-01-16T02-19-44Z
Using mc-2020-02-20
```
1,596/16,000 objects; 9 GiB in 2m55s
1,600/16,000 objects; 9 GiB in 2m54s
15,644/16,000 objects; 9 GiB in 11m16s
```

### RELEASE.2020-08-27T05-16-20Z
Using mc-2020-02-20
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

# All objects (10 buckets, 800x10kb, 400x500kb, 400x2MB)
16,000/16,000 objects; 9 GiB in 5m53s
```

# Examples
https://gist.github.com/harshavardhana/a8d1669345fd4c5d388447ab69520fcc

# Issues
2021-03: Healing a disk is taking a very long time
https://github.com/minio/minio/issues/11696
