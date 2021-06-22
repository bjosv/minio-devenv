# Findings

## standalone mode:
If single node: all 4 pods on same
Uses Deployment

## distributed mode:
On pod on each K8s node (when available)
Uses StatefulSet
StatefulSets need persistent storage
Note that the replicas value should be a minimum value of 4

## License
On April 23 2021 Harshavardhana, cofounder of Minio Inc, submitted a change which re-licensed the project from its previous
Apache V2 license to under the Affero General Public License Version 3 (AGPLv3).

Post-change issue:
https://github.com/minio/minio/issues/12143

## kind uses local-path-provisioner
/var/local-path-provisioner/pvc-51c6ae0b-4589-4e44-9267-14af3dd29fe2_default_export-minio-0/mybucket

## Healing
MinIO server now supports auto-heal.
but previouly following was needed:
`mc admin heal -r play`

## Rebalance / Heal
mc config host add ALIAS_NAME YOUR-MINIO-ENDPOINT YOUR-ACCESS-KEY YOUR-SECRET-KEY
mc admin heal --recursive ALIAS_NAME


## Network storage
There’s one caveat when using clustered network storage solution (eg. Ceph/ GlusterFS / Rook.io), the IO performance will drop by ~60% of its original performance because of data replication, but of course you don’t have to worry about data loss when the disks/nodes are missing.
https://fadhil-blog.dev/blog/rancher-local-path-provisioner/

## Enable trace
mc admin trace

## Minio releases Feb 2020 tags:
RELEASE.2020-02-07T04-56-50Z  - Minor Bug Fix Release, Breaking Change: DeleteFileBulk API h
RELEASE.2020-02-07T23-28-16Z  - Minor Bug Fix Release
RELEASE.2020-02-20T22-51-23Z  - Minor Bug Fix Release
RELEASE.2020-02-27T00-23-05Z  - Minor Bug Fix Release, Breaking Change: Use bulk locks in Multi-object delete.

## Auto healing
https://github.com/minio/minio/issues/6782

https://github.com/minio/minio/pull/10103
"It sounds like the change will automatically heal missing items up to 2 levels deep. Any files deeper than 2 levels will some random chance of being healed every 512 "cycles"."

See link for details
https://github.com/minio/minio/releases/tag/RELEASE.2020-02-27T00-23-05Z


* If you want the disk to be repopulated immediatelly, you should run `mc admin heal --recursive myminio` so that the whole instance is healed, if you didn't do this, the server would do it automatically later, when the auto-heal feature kicks in, which happens every 24 hours according to how the feature was implemented



* I have a distributed cluster with 4 nodes. After a node was taken out of service and had its volume wiped, when i start the cluster, the console shows the message that unformatted disk found,how do i repaire the broken node? help me, please, thank you!

mc admin heal -r myminio
> also gives a status banner while it is running.
