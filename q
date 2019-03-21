[33mcommit b12a9407b64849edba43e73b7bdb87031b1803e6[m
Author: mayl <mayilong1989@hotmail.com>
Date:   Sun Mar 17 19:19:35 2019 -0700

    test byfn

[33mcommit 9073d551b3dfde7b107d1cba8ee2f923a3b1aea7[m
Author: mayl <mayilong1989@hotmail.com>
Date:   Wed Mar 13 00:01:53 2019 -0700

    2019/03/13

[33mcommit cdc8cad747cc0bf8e8e83226d25a5332d3f5c7dc[m
Author: mayl <mayilong1989@hotmail.com>
Date:   Sun Feb 24 23:33:22 2019 -0800

    cla 2/25

[33mcommit 2ef34db72084cad9c30c7fa634a82fc5eeeb6080[m
Author: mayl <mayilong1989@hotmail.com>
Date:   Thu Jan 24 22:24:51 2019 -0800

    clagen init

[33mcommit 73f7e85a7d5f0eeb6a504d56b2e9bf70e8927ae8[m
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Wed Jan 23 09:55:30 2019 +0800

    [LBSBLOC-64]Intergate branch feat_encrypted_rwset to v1.4.0

[33mcommit 1988ac226a0dee6433d68e9c0ec27a45bf5b954b[m
Merge: 908739e 4047dde
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Tue Jan 22 17:07:18 2019 +0800

    Merge branch 'feat_encrypted_rwset' into release-1.4

[33mcommit 4047dde4ef6b6288b049ee8128a16d3a8f7c38bf[m
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Tue Jan 22 10:16:11 2019 +0800

    delete useless file

[33mcommit 33f36904fcf51e9f034eb7b5a986a9fa40d01e59[m
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Mon Jan 21 14:54:43 2019 +0800

    commit ignore examples and run sub-dir

[33mcommit 3ad96bc65b9da065290a0d14ef348f6091a1c603[m
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Tue Nov 6 14:41:41 2018 +0800

    modify code for pass linter verify

[33mcommit 038dbcf23db4121554e4066b78dfba23f379e79b[m
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Mon Nov 5 14:28:03 2018 +0800

    delete the obsolescent files

[33mcommit 4e7ec93b14e2e2da75fad0d93c0ffffe8f348cfc[m
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Mon Nov 5 14:23:10 2018 +0800

    format code stype with gofmt

[33mcommit 8cd4afc76551d7012ce961fc9a0e3a16b932b39a[m
Author: shuaibc1 <shuaibc1@lenovo.com>
Date:   Thu Nov 1 10:44:01 2018 +0800

    new method: GetTxIDByBlockNumTxNum

[33mcommit 5914061bc6a1770ff961a093850f9b6856f9d148[m
Author: shuaibc1 <shuaibc1@lenovo.com>
Date:   Wed Oct 31 11:08:17 2018 +0800

    PutState/GetState optionally encrypt/decrypt the state value

[33mcommit d8f08c4bc7cd4a523dca4a66a43f4c4c37671b98[m
Author: shuaibc1 <shuaibc1@lenovo.com>
Date:   Tue Oct 30 13:53:39 2018 +0800

    add GetStateVersion and GetPrivateDataVersion

[33mcommit daa2358743b185fef866672cec94c9f1f59bc105[m
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Tue Oct 23 11:34:22 2018 +0800

    merged from row_level_encrypt in fabric_v1.2

[33mcommit 949a1ea0f26429cfc2c5c1ecf9ef5ac3c0178231[m
Author: wangyh43 <wangyh43@lenovo.com>
Date:   Tue Oct 23 11:30:57 2018 +0800

    merged from row_level_encrypt in fabric_v1.2

[33mcommit 908739e0f5526cd408759d629db58b77763dac9f[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Mon Jan 21 02:39:11 2019 +0200

    [FAB-13471] lifecycle to handle multiple updates
    
    The ledger calls several HandleChaincodeDeploy for each update
    but only a single ChaincodeDeployDone after all invocations
    to HandleChaincodeDeploy were made.
    
    The current implementation only supported a single update,
    and as a result - a second ChaincodeDeployDone will get stuck
    writing to a channel.
    
    This change set makes the Lifecycle in core/cclifecycle
    to be able to handle any number of updates in a single block.
    
    Change-Id: I1d85018af398bd5cb968e42031986a999f6be444
    Signed-off-by: yacovm <yacovm@il.ibm.com>

[33mcommit f118f2bca914f74567701b26c3b2922738dc8f6f[m
Author: pama-ibm <pama@ibm.com>
Date:   Fri Jan 18 14:29:19 2019 -0500

    [FABCI-258] Fixed doc link
    
    Fixed the docs badge
    link to point to the release branch
    of the doc. Also added a link to the
    Hyperledger Fabric rocket chat
    
    Change-Id: Icbe64888496ffda046d4e2668bd0caf100875454
    Signed-off-by: pama-ibm <pama@ibm.com>

[33mcommit 782c4eea7aadf0cbdbb2f63afc0706fe81ac18ce[m
Author: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
Date:   Tue Dec 18 19:14:55 2018 +0000

    FAB-13336 Develop Apps: Chaincode namespace
    
    Change-Id: I883797b9412e411d80bd756d0e91faed964729da
    Signed-off-by: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
    (cherry picked from commit d49b0955a6e06eff463ca71c6383a470a008f245)

[33mcommit f80f06953cf8eca38c9d428533f2547a953d983a[m
Merge: 6ae588a b0331a9
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Jan 16 14:43:10 2019 +0000

    Merge "FAB-12908 Add health check for CouchDB" into release-1.4

[33mcommit 6ae588a21fc84178557f648a64aad11e6e3c57c6[m
Merge: 8f9e8c8 55a1804
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Wed Jan 16 10:07:21 2019 +0000

    Merge "[FAB-13281] Document trust relationships" into release-1.4

[33mcommit 8f9e8c8fd14761baddc07851a613ecd9ef95db2b[m
Merge: d885dec 3e6641d
Author: Yacov Manevich <yacovm@il.ibm.com>
Date:   Tue Jan 15 22:42:58 2019 +0000

    Merge "Fix intermittent test failure in acl e2e test" into release-1.4

[33mcommit b0331a935310c35df233a8163023b6a51b2956af[m
Author: Saad Karim <skarim@us.ibm.com>
Date:   Tue Dec 11 15:22:45 2018 -0500

    FAB-12908 Add health check for CouchDB
    
    Added health check for CouchDB
    
    Change-Id: If2d32b2197162b5f04997ab71ba2d292701cd0b9
    Signed-off-by: Saad Karim <skarim@us.ibm.com>
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>
    (cherry picked from commit 8768567a9c626b2feb533f3c4909d2243d44aac9)

[33mcommit d885decd6e8a39ae306bce202df3c171343a0a3c[m
Author: Artem Barger <bartem@il.ibm.com>
Date:   Fri Jan 11 02:20:21 2019 +0200

    [FAB-12065] fix TestLeaderYield flaky test
    
    Change-Id: Iaba09fe4902d421310f21378b0d8dbfb5e2ed107
    Signed-off-by: Artem Barger <bartem@il.ibm.com>
    (cherry picked from commit e9c2f4c943845f04c45c44efc3a634a36a66000f)

[33mcommit 3e6641d66b4043a7b8608167dabf310ad4c6495d[m
Author: Will Lahti <wtlahti@us.ibm.com>
Date:   Wed Jan 9 13:08:54 2019 -0500

    Fix intermittent test failure in acl e2e test
    
    FAB-13594
    
    Change-Id: I27e83ec5edef355362fa5560623786105599c2ec
    Signed-off-by: Will Lahti <wtlahti@us.ibm.com>
    (cherry picked from commit e66e67c5f5ed4ced60089848f3c18bca33199d77)

[33mcommit ba5d8dc73eae9769d6d7c359e38f2cc9156ca28b[m
Merge: bbd4273 c6ea50e
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Mon Jan 14 13:41:54 2019 +0000

    Merge "[FAB-13644] move to promhttp handler" into release-1.4

[33mcommit bbd4273dca19e830dc231d3c7eb6d2202e792116[m
Author: pama-ibm <pama@ibm.com>
Date:   Thu Jan 10 16:16:16 2019 -0500

    [FAB-13605] Updated Copyright footer
    
    to 2019.
    
    Change-Id: Ie20499d8b6e784bf418fd83d709ac6b2a7962081
    Signed-off-by: pama-ibm <pama@ibm.com>
    (cherry picked from commit 26d34aaf3340bb17348134696b873173a3c5fd54)

[33mcommit 55a1804196a36c3084c461fabc23a43928b7b2d1[m
Author: Matthias Neugschwandtner <eug@zurich.ibm.com>
Date:   Fri Dec 14 15:37:31 2018 +0100

    [FAB-13281] Document trust relationships
    
    Document the trust relationships of PaperNet.
    
    Change-Id: Ic2cea82a001833484a4220b89c8a1c936fc786ce
    Signed-off-by: Matthias Neugschwandtner <eug@zurich.ibm.com>

[33mcommit d2fbc32986efc4a8f9287461e5f80fbd359b4614[m
Author: Matthias Neugschwandtner <eug@zurich.ibm.com>
Date:   Fri Jan 11 12:19:29 2019 +0100

    [FAB-13593] Ledger synchronization in SBE tests
    
    The SBE integration tests include cases where the ledger is updated,
    waiting for the event on one peer and then checking the result on
    another peer. This patch adds synchronization that waits until both
    peers have the same ledger height, i.e. the transaction has been
    delivered to both peers.
    
    This is handled similarly in the private data integration tests.
    
    Change-Id: I6b68d3cf9cfeebf6cddd75003b3a946295fb4392
    Signed-off-by: Matthias Neugschwandtner <eug@zurich.ibm.com>
    (cherry picked from commit 0de5a321b98f60ab1060a2d1e1491880e22558f6)

[33mcommit 1b1118e53d127b1774b9345a574eda08c7811138[m
Author: Artem Barger <bartem@il.ibm.com>
Date:   Thu Dec 20 00:51:19 2018 +0200

    [FAB-13366] fix gossip state flake
    
    This commit removes time.Sleep() and instead waiting for peer to get
    connected with bootstrap peer to build membership before sending state
    message.
    
    Change-Id: If132432a710e3006fc9de0e9ea051709e6edbde6
    Signed-off-by: Artem Barger <bartem@il.ibm.com>

[33mcommit 73cc36fe38ea3db806be0df826d9d48034100bd2[m
Author: waels <wael.shama@ibm.com>
Date:   Mon Dec 10 17:33:55 2018 +0200

    [FAB-11639]: Fix data races in gossip/discovery
    
    Change-Id: I6d5c6554ffe77079ee23c735f1ea89e3654f695f
    Signed-off-by: waels <wael.shama@ibm.com>
    (cherry picked from commit e67eeb9bae8d25c7b072b7b2c6f6f0dcbc5becfe)

[33mcommit c0644669d08f6a3b38b8470393db478b9177b175[m
Author: waels <wael.shama@ibm.com>
Date:   Wed Nov 28 11:32:59 2018 +0200

    [FAB-11643]: Fix data races in gossip/state
    
    Squashed with:
    
    [FAB-13377] revert stop discovery after chanState
    
    Change-Id: I17d89c3d0d84e634c5bdb477badcdfa2205d4ac7
    Signed-off-by: waels <wael.shama@ibm.com>
    (cherry picked from commit 2fd63a497531f85937565e74e5a2c61a40fade1b)

[33mcommit 26aca2e318f903c8ceade9b6d58f1054cabd626d[m
Author: muralisr <srinivasan.muralidharan99@gmail.com>
Date:   Fri Nov 30 16:30:32 2018 -0500

    add endorser metrics
    
    This CR adds endorser metrics
       - #proposals received
       - #successful proposals
       - #proposals that failed validation
       - #proposals that failed due to tx dup.
       - #proposals that failed due to endorsement
          failures
       - #proposals that failed chaincode instantiations
       - #duration of successful proposals
    
    Metrics from other parts in the stack such as chaincode are left
    to those components.
    
    FAB-13088 #done
    
    Change-Id: I16a472540b2cd6e31c93d89cce9b5e69940d2db4
    Signed-off-by: muralisr <srinivasan.muralidharan99@gmail.com>
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>
    (cherry picked from commit 2d2cd3366bf7693e5e7545682fe1e55ed913512f)

[33mcommit c6ea50e78cb267829c79aae68022b71bef57febf[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Fri Jan 11 17:26:20 2019 -0500

    [FAB-13644] move to promhttp handler
    
    The handler we were using has been relatively recently deprecated.
    
    Change-Id: Ie32b3b5a2bb50a37a134a94c3dba7d3c6b88b47f
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit ddff870379eac275b2fba8d766b4cf382e1095a5[m
Merge: 8faadaf 80e1b23
Author: Yacov Manevich <yacovm@il.ibm.com>
Date:   Fri Jan 11 21:30:56 2019 +0000

    Merge "[FAB-13598] remove grpc_start_time from logs" into release-1.4

[33mcommit 8faadaf375830c10d73907530eca65a5df1a490f[m
Merge: ec8d9ea 481ae97
Author: Yacov Manevich <yacovm@il.ibm.com>
Date:   Fri Jan 11 20:46:52 2019 +0000

    Merge "[FAB-13602] Fix time.Timer leak in gossip handshake" into release-1.4

[33mcommit ec8d9eaed9db7cfc5d2f6d15248163ae3a0a1282[m
Merge: 56dfcad 880709a
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Fri Jan 11 16:11:02 2019 +0000

    Merge "[FAB-13591] reduce scope of rlock around observer" into release-1.4

[33mcommit 56dfcadc9a68782b6b2cd311a78c8361d72d0a03[m
Merge: 6a82f9d e51851a
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Fri Jan 11 16:10:57 2019 +0000

    Merge "[FAB-13237] metrics for log records" into release-1.4

[33mcommit 80e1b23d30da2b4ef05b14834ebdd3c7b8a77aab[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Thu Jan 10 09:10:50 2019 -0500

    [FAB-13598] remove grpc_start_time from logs
    
    Since we have the duration in the log record and the log record includes
    a time stamp, determining the start time is trivial relative to the
    amount of text produced when formatting a time stamp.
    
    Change-Id: I37e926ffd8402400c58e3bffabc217efc749717c
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit 6a82f9d5ef7af4eff82d405944253167904516cc[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Thu Jan 10 15:39:10 2019 -0500

    [FAB-13627] Add LTS to v1.4 What's New doc
    
    Add long term support policy
    to v1.4 What's New doc.
    
    Change-Id: Ic375c9db3a10b79590c252bde5d16a84a96ff15a
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit 51b410c9b22e7f9d1b3ceb12df2598dd6d2eca8c[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Jan 9 18:38:56 2019 -0500

    [FAB-13556] Prepare fabric for next rel (v1.4.1)
    
    Change-Id: Iffef1bff4f3fa2b972c6c54a4b12540bd93255e1
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit 481ae9784a8f71b9e01d574473b28cfc0a2bda5d[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Thu Jan 10 00:03:28 2019 +0200

    [FAB-13602] Fix time.Timer leak in gossip handshake
    
    When authenticating a remote peer, a connection message
    is expected to be received from it.
    
    There is a time.Timer used, which isn't closed, and therefore
    there is a resource leak of a time.Timer.
    
    This change set changes this to a time.After.
    
    Change-Id: I4b4dce2af2854ccd1e061296daa16b0c5f42c525
    Signed-off-by: yacovm <yacovm@il.ibm.com>

[33mcommit 8f9c5e5a7967fedf78ceb754aeea0d881cb1e600[m
Merge: d700b43 ae96bf5
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Jan 9 17:40:02 2019 +0000

    Merge "[FAB-13347] throttle grpc concurrency" into release-1.4

[33mcommit d700b43476e803c864c48021e63a78543b60e17e[m
Merge: 4fd7013 69bfba8
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Jan 9 15:47:59 2019 +0000

    Merge "[FAB-13555] Release fabric v1.4.0" into release-1.4

[33mcommit 880709a6baa99800ea8c3da2cb9b00ab6b73ce4e[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Jan 9 10:09:01 2019 -0500

    [FAB-13591] reduce scope of rlock around observer
    
    Change-Id: I7ea50a583b0af311b02366a1e5124dd9faedf380
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit e51851a3c0e5b878e894151ed2c057e81cc8707f[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Fri Jan 4 14:25:28 2019 -0500

    [FAB-13237] metrics for log records
    
    Created counters to track the number of log records checked and
    written. (filtered = checked - written)
    
    Change-Id: I1b4dcbdc636891e8deca41440c40d58415edc438
    Signed-off-by: Saad Karim <skarim@us.ibm.com>
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>
    (cherry picked from commit 312f11343942d312510027a8a4787cf3be98fce1)

[33mcommit ae96bf50153a36b18bf21eccc63d9abac8e9bcc6[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Jan 8 15:00:40 2019 -0500

    [FAB-13347] throttle grpc concurrency
    
    Change-Id: I912f3a28525c6583caee451adfd25578e7c299ad
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>
    (cherry picked from commit 33800edf869bb3127992a14ec01e4890365f641a)

[33mcommit 4fd7013ee15cf88130244f030a2f3a20f3ba031c[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Jan 8 13:38:46 2019 -0500

    [FAB-13347] introduce counting semaphore
    
    The counting semaphore will be used as part of throttling
    implementation.
    
    Change-Id: Ie8741414a6b25373072a611a55620519051489cc
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>
    (cherry picked from commit f975549503d76c276adfad52c1061901526bb410)

[33mcommit 69bfba80a42a6948161bbcd01c775b9e8911529a[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Tue Jan 8 12:16:38 2019 -0500

    [FAB-13555] Release fabric v1.4.0
    
    Change-Id: I8e568ab2088e68df2e8c804dfef2bce5a311f6fe
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit b43fcc774476a1d931fe8ff782eceb4ab61196a2[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Tue Jan 8 00:57:08 2019 -0500

    [FAB-12056] Pvt data tutorial to use transient
    
    fabric-samples has been udpated to pass private data
    using the transient field. This CR updates the
    private data tutorial to be in sync with the sample.
    
    Change-Id: Id610cf7dfd528f3f294227c9934b6c03a3c4236b
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit a06f4eaa271798aeaedbcc1ea7ff9c966117581c[m
Merge: 491df44 24c409a
Author: Yacov Manevich <yacovm@il.ibm.com>
Date:   Sat Jan 5 18:16:09 2019 +0000

    Merge "[BE-510] Policies documentation typo fix" into release-1.4

[33mcommit 491df44791aada8024094dfbd071f9c17d269655[m
Author: Amol Pednekar <amol_pednekar@persistent.com>
Date:   Fri Jan 4 12:55:14 2019 +0530

    [FAB-13357] Fixed error in fabric gossip documentation
    
    In fabric's Gossip data dissemination protocol documentation,
    under "Static leader election", point #2 had wrong variables,
    both being CORE_PEER_GOSSIP_USELEADERELECTION. The second variable
    has been updated to CORE_PEER_GOSSIP_ORGLEADER
    
    Change-Id: I1be381ed3da1face2b1fc2210a5d783c7693e8f5
    Signed-off-by: Amol Pednekar <amol_pednekar@persistent.com>

[33mcommit 24c409ac0b0792ad867149b0319a35147b4c60f3[m
Author: berendeanicolae <berendeanicolae@gmail.com>
Date:   Fri Jan 4 15:31:53 2019 +0200

    [BE-510] Policies documentation typo fix
    
    Fix typo in documentation for policies. The complex example used mspP1 and the
    explanation mentioned mspP0.
    
    Change-Id: I70d61339c8a2896b10d9993040d1be19dee5e233
    Signed-off-by: berendeanicolae <berendeanicolae@gmail.com>

[33mcommit 93b2575a85c9f5736b5696270b5ae5947dbd3dc1[m
Merge: a83e1ca 7203735
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Fri Jan 4 22:29:54 2019 +0000

    Merge "[FAB-12995] Add new functions to the example cc" into release-1.4

[33mcommit a83e1ca58b63e488599b6ba31afabea1b1f1bf84[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Thu Dec 20 13:56:48 2018 -0500

    [FAB-13411] fix flake in TestServerInterceptors
    
    The server stream interceptor added by TestServerInterceptors returns an
    error that causes the server side of the stream to terminate. That means
    the client side send may get an io.EOF from Send if the termination of
    the server side of the stream was detected.
    
    Change-Id: I237f0aad6daa2878a987509e35074a16b689feba
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>
    (cherry picked from commit 6e56e6e30274579b1895d84980a4bd5b8950a52e)

[33mcommit 7203735ca299e568c4ff585ea40bc3d3a630bc85[m
Author: ratnakar <asara.ratnakar@gmail.com>
Date:   Thu Jan 3 15:04:49 2019 -0500

    [FAB-12995] Add new functions to the example cc
    
    Change-Id: Ia5f062a72c5edcbb57ce44249a15044625b8f92c
    Signed-off-by: ratnakar <asara.ratnakar@gmail.com>

[33mcommit 7bb9b7e5f42c10b3adf28523c9d7531311d592e9[m
Merge: 3c0d63c 155120d
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Jan 2 16:35:26 2019 +0000

    Merge "Increase timeout for TestHaltBeforeTimeout" into release-1.4

[33mcommit 3c0d63c3db8cb84a57ff49fd984712a5393f5ed6[m
Author: Gari Singh <gari.r.singh@gmail.com>
Date:   Wed Jan 2 09:26:29 2019 -0500

    Explicitly set ext key usage for CA
    
    The generated CA certificates currently have
    contain anyExtendedKeyUsage in their
    Extended Key Usage attributes.  This is
    actually not allowed and is now enforced by
    openssl 1.1 and later.
    
    This change explicitly adds only ClientAuth
    and ServerAuth to the CA's Extended Key
    Usage attributes.
    
    FAB-13439 #done
    
    Change-Id: Ia2586563bd46c2978704999d1d9307d110bbcc98
    Signed-off-by: Gari Singh <gari.r.singh@gmail.com>

[33mcommit 155120dcdf2639059fbcb0c5fce4a15f5169635f[m
Author: Will Lahti <wtlahti@us.ibm.com>
Date:   Thu Dec 27 12:49:17 2018 -0500

    Increase timeout for TestHaltBeforeTimeout
    
    Testing in CI. Will update commit message if
    this fixes the flake.
    
    FAB-13375
    
    Change-Id: I3bad8747cd2c5d47f79ceddaf78f1289c35f7fe1
    Signed-off-by: Will Lahti <wtlahti@us.ibm.com>
    (cherry picked from commit 58429827b6fa4300c37dd7b7813c9c44c11aee57)

[33mcommit 7d09ca647b1b95fa961eacf480bf8cffd7c1cbcb[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Tue Jan 1 12:24:35 2019 -0500

    [FAB-13463] Document vendoring - part2
    
    Fix vendoring link.
    
    Change-Id: Ie21ec4fc9e3dfc171948ff5c633da7ae8dd6bf17
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit b17d88ff0ffaefe45da99866a006853d94cbddcc[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Mon Dec 31 13:31:00 2018 -0500

    [FAB-13463] Document vendoring for shim extensions
    
    Add documentation to clarify how to vendor
    shim extensions into chaincode.
    
    Change-Id: Ic568e00c19cd512cfd2160b109ac9090c8a38901
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit 99959b9f8d5dbeb4e727b0f21e102aa19acf8402[m
Author: kuro1 <412681778@qq.com>
Date:   Thu Dec 20 10:58:48 2018 +0800

    [FAB-13381] Update dev mode documentation
    
    - Update a chaincode path and keep it a bit vague about
      the exact location of the config file.
    - Add an environment variable when start peer dev-mode
    
    Change-Id: Ie5ef5f1399efa5960373b984223209132c72cd96
    Signed-off-by: kuro1 <412681778@qq.com>
    (cherry picked from commit 274ce1fbe1c9de4eaf3b5fb9df7a8cb4c280d022)

[33mcommit 9cd9fce9492a31b5666ba4defd83830d913e67a2[m
Author: Angelo De Caro <adc@zurich.ibm.com>
Date:   Fri Dec 21 14:26:21 2018 +0700

    [FAB-13351] Test Robustification
    
    This change-set enhance the tests to make sure
    that a credential is invalidate always successfuly.
    
    Change-Id: Id6d11bc735f0b03a07e97e99a57ec41b4ea26bb8
    Signed-off-by: Angelo De Caro <adc@zurich.ibm.com>
    (cherry picked from commit 419397e47d6fc595fd3ff3154f6688151a949b61)

[33mcommit 0311c83af4ca5ea223bf343eddb79a76adfe6f49[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Thu Dec 20 14:25:40 2018 -0500

    [FAB-13391] Prepare for next release (1.4.0)
    
    Change-Id: I001d80286f1e58ea29d62a6af6433165471e5e8b
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit b87ec8086e8bf65ca8a6bdde4001247413da93ce[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Thu Dec 20 09:01:09 2018 -0500

    [FAB-13390] Release fabric v1.4.0-rc2
    
    Change-Id: Icaa3581fd491874ac85c960280ed10ab676b6c82
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit f87c3f9635d93d508df5698faf5984b174341672[m
Merge: 5c54bf2 47317b2
Author: Manish Sethi <manish.sethi@gmail.com>
Date:   Wed Dec 19 22:26:53 2018 +0000

    Merge "recon: add debug logs in ledger" into release-1.4

[33mcommit 47317b23f91a8ea0495c497f46199a1931a3e5f2[m
Author: Senthil Nathan N <cendhu@gmail.com>
Date:   Wed Dec 19 11:35:49 2018 +0530

    recon: add debug logs in ledger
    
    FAB-13356 #done
    
    Change-Id: I8359ae52b6a0c0bcf203c6f4be4faa27563f36ad
    Signed-off-by: senthil <cendhu@gmail.com>
    Signed-off-by: manish <manish.sethi@gmail.com>

[33mcommit 5c54bf25de4df881144db3b7eecda25b59307e5f[m
Author: Senthil Nathan N <cendhu@gmail.com>
Date:   Wed Dec 19 19:45:20 2018 +0530

    opt: lscc state cache in couchDB
    
    FAB-13368 #done
    
    Change-Id: Ib6404b104aa5305aa09f64ddfd9007f7770b700e
    Signed-off-by: senthil <cendhu@gmail.com>

[33mcommit 4fab33e3ae181799ce840c0587438023a17bfea6[m
Merge: 074aa96 e9e5ca6
Author: Gari Singh <gari.r.singh@gmail.com>
Date:   Wed Dec 19 11:28:01 2018 +0000

    Merge "[FAB-13025] generate rst metric tables" into release-1.4

[33mcommit 074aa9667a624006ab1a2ce6417c1348c655ca97[m
Merge: d1a8e80 f91e87c
Author: Gari Singh <gari.r.singh@gmail.com>
Date:   Wed Dec 19 11:27:37 2018 +0000

    Merge "Add meter with fabric and go version" into release-1.4

[33mcommit d1a8e80ed23e464c78bbe9dca8ad8ff68a64ebf9[m
Merge: f671c9a ecb8e48
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 19 03:36:06 2018 +0000

    Merge "FAB-13271 Commercial Paper Updates" into release-1.4

[33mcommit f671c9a9a460fa163890d7c135f55992a1f5d9eb[m
Merge: a467c65 863f989
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 19 03:32:46 2018 +0000

    Merge "recon: at a time, only 1 func. can use the Cache" into release-1.4

[33mcommit ecb8e48432119de29d026d485ab50c8f8ae1ae75[m
Author: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
Date:   Mon Dec 17 14:32:57 2018 +0000

    FAB-13271 Commercial Paper Updates
    
    Change-Id: I82cb88af9837a1725e45489d9b6db2e61efb2bc1
    Signed-off-by: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
    (cherry picked from commit 5c887e5e3666eba65c55245e993f814f566c8437)

[33mcommit a467c654d293abc0cbdbabec02ef4bbf11dca80f[m
Merge: 4f42c37 85802b9
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Tue Dec 18 18:25:44 2018 +0000

    Merge "FAB-12978 Develop Apps: Connection options topic" into release-1.4

[33mcommit 4f42c37710fe0d5624d3093fe80a4f367634225a[m
Merge: c921e37 728f4eb
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Tue Dec 18 18:17:19 2018 +0000

    Merge "[FAB-13327] Architecture Explained clarification" into release-1.4

[33mcommit c921e37ee8fa000cf0746fadea9868fee90fe11c[m
Merge: df845e0 283bb61
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Tue Dec 18 18:16:51 2018 +0000

    Merge "FAB-12934 Develop Apps: Connection profile topic" into release-1.4

[33mcommit 85802b93557e2011d6afe682027b26a98a585b8b[m
Author: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
Date:   Fri Nov 23 07:30:33 2018 +0000

    FAB-12978 Develop Apps: Connection options topic
    
    Change-Id: Ib1f0ded63b01b4b8ba4a409fee853b8c58e697ed
    Signed-off-by: Anthony O'Dowd <a_o-dowd@uk.ibm.com>

[33mcommit e9e5ca6d4415ee0323c10343260066650136e80a[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Mon Dec 3 16:27:40 2018 -0500

    [FAB-13025] generate rst metric tables
    
    Use option definitions in the source tree to generate restructured text
    tables to be included in doc.
    
    Change-Id: I0243581ce10875cf1bbebe2b3e6765f2b91150e6
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit f91e87c959517531a1b76792bcc43a5ee7c6ab89[m
Author: Will Lahti <wtlahti@us.ibm.com>
Date:   Tue Dec 11 15:38:20 2018 -0500

    Add meter with fabric and go version
    
    FAB-13190 #done
    
    Change-Id: I038119a3f5e7c70164093788c4186ad6f50aa625
    Signed-off-by: Will Lahti <wtlahti@us.ibm.com>
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>
    (cherry picked from commit 969faf9f70ce89cfa430d333a8c5f2704d8e11c7)

[33mcommit 863f989a85a7f93eec0d6111cc010b5c0c41a847[m
Author: Senthil Nathan N <cendhu@gmail.com>
Date:   Tue Dec 18 17:38:23 2018 +0530

    recon: at a time, only 1 func. can use the Cache
    
    Among ValidateAndPrepare(), PrepareExpiringKeys(), and
    RemoveStaleAndCommitPvtDataOfOldBlocks(), we can allow only one
    function to execute at a time. The reason is that each function calls
    LoadCommittedVersions() which would clear the existing entries in the
    transient buffer/cache and load new entries (such a transient buffer/cache
    is not applicable for the golevelDB). As a result, these three functions
    can interleave and nullify the optimization provided by the bulk read API.
    Once the ledger cache (FAB-103) is introduced and existing
    LoadCommittedVersions() is refactored to return a map, we can allow
    these three functions to execute parallely.
    
    This CR ensures that these functions are not executed parallely.
    
    Further, when we commit the old pvtdata, some of these pvtdata might be
    expiring at the next block commit. As the toPurgeList might have been already
    constructed by the last Commit() call, we need to update the toPurgeList
    with eliglble old pvtdata if it is expiring. Till now, we created new
    toPurgeList by calling PrepareExpiringKeys() from
    RemoveStaleAndCommitPvtDataOfOldBlocks().
    
    This CR instead updates the existing toPurgeList rather can calling
    PrepareExpiringKeys(). This is because, we don't want to execute these
    functions in parallel. Also, it is optimal to update the existing
    expiring keys list rather than creating a new one especially when no
    new regular blocks are being committed for a longer duration and a lot
    of old pvtData are being committed.
    
    FAB-13328 #done
    FAB-13329 #done
    
    Change-Id: Ia715bd1a376c6ce19caa7af6f6510e0cf6f5a1dd
    Signed-off-by: senthil <cendhu@gmail.com>

[33mcommit 283bb617998a8a977f1a0d58a21c80d50012fc76[m
Author: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
Date:   Mon Dec 10 08:11:18 2018 +0000

    FAB-12934 Develop Apps: Connection profile topic
    
    Change-Id: I5b49c6b86f2d0efd258b249c489b3583ed2949de
    Signed-off-by: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
    (cherry picked from commit dd455f21088f777007ad9f01ef0986a26a42a28e)

[33mcommit 728f4eba340cc476efaf2c51c0b04c8b8a24337e[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Mon Dec 17 22:15:10 2018 -0500

    [FAB-13327] Architecture Explained clarification
    
    The Architecture Explained document is taken from
    the initial Fabric architectural proposal.
    
    This CR clarifies it as an Architecture Origin
    document that does not 100% represent Hyperledger
    Fabric as implemented.
    
    Change-Id: I4da44e053d55992458e8b7a292fd0f90c367ba49
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>
    (cherry picked from commit 4fe73d28512f1c87b0573698dbddef54bb67b5cd)

[33mcommit df845e0a0a457a09724ac086d7bdad81d7c19c40[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Tue Dec 18 11:07:50 2018 -0500

    [FAB-13340] Fix reconciliation CouchDB bulk load
    
    Fix bug so that bulk load cache becomes effective
    during commit of reconciled private data.
    
    Change-Id: I3ff8f8e8593cae2b5c88fd594062035fb58d5706
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit cfdb44df0d1ebfa7821026fc8e4fab24bb8d54e6[m
Author: Artem Barger <bartem@il.ibm.com>
Date:   Mon Dec 17 00:45:38 2018 +0200

    [FAB-13303] increase reconciliation test coverage
    
    Change-Id: Ifedef26a3abc30a5dfcdea597b09df8746593999
    Signed-off-by: Artem Barger <bartem@il.ibm.com>

[33mcommit 44325b46b35c9a0a00da2a010fb76dacfe295857[m
Author: Artem Barger <bartem@il.ibm.com>
Date:   Mon Dec 17 00:34:44 2018 +0200

    [FAB-13302] check for err before report mismatches
    
    This commit changes order of the reconciliation item, first checks
    whenever commit has succeeded and only then report mismatches.
    
    Change-Id: If4b91bc8afc9fe1ed282f8ef40ebfc420496927b
    Signed-off-by: Artem Barger <bartem@il.ibm.com>

[33mcommit 1954f364e11694e2fe83d5317c9251cd8fc540a4[m
Author: Artem Barger <bartem@il.ibm.com>
Date:   Mon Dec 17 00:31:24 2018 +0200

    [FAB-13301] use generated mocks in pvt data tests
    
    This commit changes pvt data tests to use auto generated mocks for
    committer interface.
    
    Change-Id: I05190e4532511966c62a371dc6abb7ebbf68de33
    Signed-off-by: Artem Barger <bartem@il.ibm.com>

[33mcommit e5b3481cb830f936eca0c9033c940bac962ec94e[m
Author: wenjian3 <wenjianq@gmail.com>
Date:   Fri Dec 14 11:36:35 2018 -0500

    FAB-13283 Update commercial paper tutorial
    
    Related to:
    https://gerrit.hyperledger.org/r/#/c/28178/
    
    - document "port number" can be passed to monitordocker.sh
      if the default port is already in use
    
    Change-Id: I59c2db715d5cf238f947c867dfd881a7839dac71
    Signed-off-by: Wenjian Qiao <wenjianq@gmail.com>
    (cherry picked from commit 5f3dc6959ecf018eb372147d5c70fa2848abd76f)

[33mcommit fb7ed549b5faba8dd1bf8fddfd879211b32b36de[m
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Fri Dec 14 20:08:14 2018 +0100

    [FAB-13292] fix version check in bootstrap.sh
    
    The check for the specified version string should be augmented by wrapping
    the command line argument in quotation marks: if no version is provided
    (which is allowed), the test fails because the non-quoted variable disappears
    and the test bash built-in has two illegal switches next to one another,
    leading to
    
    bash: line 181: [: too many arguments
    
    Change-Id: I6ca1d654ef38d6ea0fabcd3ca24c0d861321d489
    Signed-off-by: Alessandro Sorniotti <ale.linux@sopit.net>
    (cherry picked from commit 56c855d73529cd81bec25e7ee33d10ac937a8c9c)

[33mcommit 158392da09ad5d6e24b82a0b12490c2f8fa05e8c[m
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Fri Dec 14 16:31:29 2018 +0100

    [FAB-13282] Clarify version of node.js
    
    The current wording in the prerequisite document is unclear with regards to
    which version of node.js is required.
    
    Change-Id: I2529ad712d7f4f3a371967fe37e1e5198d2a1614
    Signed-off-by: Alessandro Sorniotti <ale.linux@sopit.net>
    (cherry picked from commit 5298e23154b1489efa355215721b8929c4950a4f)

[33mcommit ba37eff834dbb7ff9a363ca131350aad99d7a984[m
Author: Artem Barger <bartem@il.ibm.com>
Date:   Fri Dec 14 17:53:33 2018 +0100

    [FAB-13269] keep reconcile pvt in one pass
    
    This commit contains following changes:
    
    1. As long as there are more private data to reconcile continue fetching
    missing pieces from the other peers, instead of going to sleep each time
    we finish pull batch.
    2. Add UT to actually check we fetching in
    single pass.
    3. Reduce default sleep interval from 5m to 1m.
    
    Change-Id:I58b3f21adaa2305817f6cb144e3e7633f6b4ac92
    Signed-off-by: Artem Barger <bartem@il.ibm.com>
    (cherry picked from commit cc09957c279358422547095bdbef4b87c4a2a378)

[33mcommit 3bcdac59b19e004f05f28f0df12243506f8d49c7[m
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Fri Dec 14 18:32:28 2018 +0100

    [FAB-13288] BYFN manual step fixes
    
    The manual steps in the BYFN doc were failing when a query was issued
    against a peer that hadn't joined the channel. This was fixed, along
    with other minor typos.
    
    Change-Id: I7732784d0dd800fd2f2ff0edf34f4989e26c3cab
    Signed-off-by: Alessandro Sorniotti <ale.linux@sopit.net>

[33mcommit 02af1941a2f156930908d0dd4ceb22d319453fec[m
Author: wenjian3 <wenjianq@gmail.com>
Date:   Thu Dec 13 20:38:23 2018 -0500

    [FAB-13277] Fix typos in doc
    
    Change-Id: I4d1302e8a4a01a97623e204f65bd546356bc0147
    Signed-off-by: Wenjian Qiao <wenjianq@gmail.com>

[33mcommit bd3fbd33f895c87d38f07f25172c9b05774eaec2[m
Merge: 1ffa0fd afb2e8d
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Fri Dec 14 15:09:10 2018 +0000

    Merge "[FAB-11734] BYFN endorsement policy" into release-1.4

[33mcommit afb2e8d42998fe1fdfe8713d163244c37c19e6b3[m
Author: pama-ibm <pama@ibm.com>
Date:   Thu Dec 13 17:11:07 2018 -0500

    [FAB-11734] BYFN endorsement policy
    
    Added a step to install the cc
    on peer0.Org2 so the endorsement
    works
    
    Change-Id: I3be42c23a0232b82e60443ecae0056c7c894e170
    Signed-off-by: pama-ibm <pama@ibm.com>

[33mcommit 1ffa0fdc62fc5af9414caa11ecbf99f5cf04b4b8[m
Merge: 2cda36c da5de24
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Fri Dec 14 04:58:54 2018 +0000

    Merge "[FAB-13224] fix blocksprovider unit test flakes" into release-1.4

[33mcommit da5de24d473ac95bd67d3e0d752873272edf0d08[m
Author: Artem Barger <bartem@il.ibm.com>
Date:   Tue Dec 11 00:44:09 2018 +0200

    [FAB-13224] fix blocksprovider unit test flakes
    
    This commit takes care of intermittent CI failure due to flaky tests of
    blocksprovider_test.go. Moves assertion of recieved blocks versus added
    to payload buffer only after blocks provider has been stopped.
    
    Change-Id: Ic0cc4dcb86715288455e547efd86c3e28b96169c
    Signed-off-by: Artem Barger <bartem@il.ibm.com>
    (cherry picked from commit da45fdd89e4a1f381797cdcf5be50c69fe1af9ac)

[33mcommit 2cda36cdc6b66030291b50afee394d548fb14436[m
Author: pama-ibm <pama@ibm.com>
Date:   Thu Dec 13 14:57:15 2018 -0500

    [FAB-13270] fix mkdir cmd
    
    change mkdir -P  to mkdir -p
    
    Change-Id: I60470951e86cf1204f55d0ee905ec9ae74e4582b
    Signed-off-by: pama-ibm <pama@ibm.com>

[33mcommit 087ea413edf78e29d0d0c6e3ebde954965b45aaf[m
Merge: 862db19 ccf843c
Author: Artem Barger <bartem@il.ibm.com>
Date:   Thu Dec 13 09:12:25 2018 +0000

    Merge "[FAB-11608] Gossip: Optimize alive message verification" into release-1.4

[33mcommit 862db196bd4a5c53d8c2c494f751b8f725fc0c41[m
Merge: c10e00c a00a80b
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Thu Dec 13 03:30:23 2018 +0000

    Merge "[FAB-13251] Add nil check in ToGossipMessage" into release-1.4

[33mcommit c10e00c7cd7d031b654dcb115db3267b5b5dcc39[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 12 14:54:32 2018 -0500

    [FAB-13256] Fix What's New link to release notes
    
    Point to rc1 release notes for now.
    
    Change-Id: I3164a500eca2442efccd917fda92ae69d739575a
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit c5abc91c5f0346829a76030fa85c363aa32cae79[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 12 12:12:15 2018 -0500

    [FAB-13253] Improve SDK compatibility doc
    
    Clarify support statement.
    
    Change-Id: I248b96a417473b27bd1ed30cb8face2f949401c2
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit a00a80b6b959e7406a1c84d94c3300e205e3b41e[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Wed Dec 12 18:23:02 2018 +0200

    [FAB-13251] Add nil check in ToGossipMessage
    
    Change-Id: I4b7dac64775107d88b0807a4e5f058222e6b8d2f
    Signed-off-by: yacovm <yacovm@il.ibm.com>

[33mcommit ccf843c8c7a7b4929849d3a6dc33ea4860f4b351[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Wed Dec 5 18:23:32 2018 +0200

    [FAB-11608] Gossip: Optimize alive message verification
    
    Alive messages are currently verified in 2 places:
    1) Gossip routing layer, before forwarded to the gossip membership layer
    2) Gossip membership layer
    
    This causes redundant message verification which induces CPU overhead due
    to un-needed ECDSA signature verification.
    
    This change set:
    
    1) Removes alive message verification in the routing layer
    2) Solidifies the logic of the verification in the membership layer
    3) Adds a unit test that ensures that:
       - Alive messages are verified "once"
       - Alive messages from membership responses are verified
         if they are fresher than the messages received directly
         by gossip so far.
       - Alive messages from membership requests are always verified
         since they trigger sending membership responses based on
         the membership request's alive message.
    
    Change-Id: I70f8c0f1acfdf52e2107fd71ad55a3d002a392ec
    Signed-off-by: yacovm <yacovm@il.ibm.com>
    (cherry picked from commit f4067cb7fec2fb3c46372e49a19de323d709d4e8)

[33mcommit 681dcfc108c63471b98098bf98655daf7de446df[m
Merge: d50b3a8 05d99ad
Author: Manish Sethi <manish.sethi@gmail.com>
Date:   Wed Dec 12 16:47:03 2018 +0000

    Merge "regenerate ledger testdata using release-1.1" into release-1.4

[33mcommit d50b3a8565801b57d518f24133f5ab7e921cee69[m
Merge: 0b7aefd 3cbc5cf
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Wed Dec 12 16:39:03 2018 +0000

    Merge "Remove +lifecycle enable parameter from core.yaml" into release-1.4

[33mcommit 3cbc5cff4399f8e9451b56eddfd5c3489a649179[m
Author: Will Lahti <wtlahti@us.ibm.com>
Date:   Wed Dec 12 10:37:05 2018 -0500

    Remove +lifecycle enable parameter from core.yaml
    
    This parameter was accidentally left in the codebase
    and should be removed to avoid any confusion regarding
    lifecycle.
    
    FAB-13249 #done
    
    Change-Id: I632591004e943f1294780b123049ece1769e3d67
    Signed-off-by: Will Lahti <wtlahti@us.ibm.com>

[33mcommit 0b7aefdffda3d66e11c7159b3131d469fbcba962[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Dec 11 14:27:40 2018 -0500

    [FAB-13239] terminate container streaming output loop
    
    When vm.docker.attachStdout is enabled, output from chaincode containers
    is written to the peer log. A recent change went in that broke the error
    handling behavior and resulted in a tight loop of reading from a closed
    reader and issuing an error message. This change fixes the error
    handling and back-fills test.
    
    Change-Id: Icda853dba90b873f2fbddde990c5f61f7834f1c9
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit 05d99ad1726e94d7d73f2b30374179ea6958c260[m
Author: senthil <cendhu@gmail.com>
Date:   Thu Nov 22 14:38:30 2018 +0530

    regenerate ledger testdata using release-1.1
    
    we need to regenerate the testdata using fabric-release v1.1 with the full
    definition of collection configs including MembersOrgsPolicy. This is necessary
    as one of the stateUpdate listener (i.e., collElgNotifier) compares the existing
    MembersOrgsPolicy with the new MembersOrgsPolicy (during upgrade tx) to perform
    certain tasks in the ledger. As the existing test data does not contain
    MembersOrgsPolicy in the used collectionConfigPkg, it leads to nil pointer
    exception.
    
    This CR creates a new testdata with fabric-release v1.1 using
    https://gerrit.hyperledger.org/r/#/c/22749/ (which is updated to include
    the MembersOrgsPolicy in the collectionConfigPkg)
    
    This CR also enables the v11 ledger test.
    
    FAB-12969 #done
    
    Change-Id: I3d8353145ca7f753f82805b7802d7a68fde9d638
    Signed-off-by: senthil <cendhu@gmail.com>
    Signed-off-by: manish <manish.sethi@gmail.com>
    (cherry picked from commit aefa370d3f799aaac7178ff7cd4030b8673b9d2d)

[33mcommit e485f77ec5599d5632e4205daea09a522790ae83[m
Author: manish <manish.sethi@gmail.com>
Date:   Tue Dec 11 19:29:56 2018 -0500

    Fix bug in decoding missingdatakey
    
    FAB-13240 #done
    
    Change-Id: I7ec5191a20743c99e992e7ce4b3c2a6117087144
    Signed-off-by: manish <manish.sethi@gmail.com>

[33mcommit 5bcbbb5cc4ecc038a40854e103b703a515646bba[m
Author: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
Date:   Mon Dec 10 10:34:35 2018 +0000

    [FAB-13024] Update fabcar doc
    
    Updates for fabric-network programming model.
    
    Change-Id: I45850245802619ace3228af77dda0d22a3b408cd
    Signed-off-by: Anthony O'Dowd <a_o-dowd@uk.ibm.com>
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit 0b07d53edbeea59889dbe973d3235a597d8fc83c[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Fri Dec 7 19:00:45 2018 -0500

    [FAB-12870] fix timing flake in grpclogging tests
    
    Ensure the client stream is setup before the server returns an error.
    
    Change-Id: Id5062f2947d44588b32127afc913c36ff238f834
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit 11fa186a0b468d902d956a79945f73ed97477302[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Tue Dec 11 00:32:46 2018 -0500

    [FAB-13114] Prepare for next rel (1.4.0 on release-1.4)
    
    Change-Id: I8d95e56adc24086d347c41f83e741d97128b858c
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit eca1b14b7e3453a5d32296af79cc7bad10c7673b[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Mon Dec 10 13:10:43 2018 -0500

    [FAB-13113] Release fabric v1.4.0-rc1
    
    Change-Id: I66fc1f9ef8f9820cb8b990f3cf18d845a5a66343
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>

[33mcommit 6dfd1feee80ce4e7582b22bcde7e56c42ea6ffb6[m
Merge: 942052f d3fda51
Author: Jonathan Levi (HACERA) <jonathan@hacera.com>
Date:   Mon Dec 10 16:34:20 2018 +0000

    Merge "FAB-13155 Update multiarch script"

[33mcommit 942052f3f8e76c047aaf82a193c291f91a1ce2f9[m
Author: joe-alewine <Joe.Alewine@ibm.com>
Date:   Mon Dec 10 10:36:07 2018 -0500

    [FAB-13214] Change link in upgrade doc
    
    Passage should point to release
    notes, not the operations
    service doc
    
    Change-Id: I4b161ec7e3c98eff481ee4718454882292eaea49
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit e58ab1249ad8a405de99baf5508b89188c84313b[m
Author: joe-alewine <Joe.Alewine@ibm.com>
Date:   Mon Dec 10 10:23:20 2018 -0500

    [FAB-13213] Add CA note to operations
    
    It's important to make it explicit
    that the TLS CA for the operations
    service should be separate and
    dedicated to it.
    
    Change-Id: Ic738a71c70a4e4b0913c9899d357ea925562b6fb
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit 39b57cd0cb8767520a15d8a5433c96317802a962[m
Merge: 1bd3f3b 2b6fcf4
Author: Kostas Christidis <kostas@gmail.com>
Date:   Mon Dec 10 15:06:25 2018 +0000

    Merge changes from topics 'FAB-13056', 'FAB-13055', 'FAB-13054'
    
    * changes:
      [FAB-13056] Onboarding: systemchannel from boot block
      [FAB-13055] Don't close LedgerFactory in onboarding
      [FAB-13054] Prevent int overflow in IsReplicationNeeded

[33mcommit 1bd3f3be77061f8e56a7342c257839972b2e2382[m
Merge: e48e5b6 7979d31
Author: Kostas Christidis <kostas@gmail.com>
Date:   Mon Dec 10 15:04:30 2018 +0000

    Merge "[FAB-13210] Remove unused code in orderer/.../main.go"

[33mcommit e48e5b6454160df18945afed9a4ea03b3b95369f[m
Merge: 081166a 1433a99
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Mon Dec 10 14:26:29 2018 +0000

    Merge "[FAB-13021] What's new in v1.4"

[33mcommit 7979d319b181c29c8287d41420810d413cf15f38[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Mon Dec 10 15:53:38 2018 +0200

    [FAB-13210] Remove unused code in orderer/.../main.go
    
    This change set removes unused code in
    the file orderer/common/server/main.go
    
    Change-Id: I5b5b3fe960633c8400008ba51a71e2c63344bdcf
    Signed-off-by: yacovm <yacovm@il.ibm.com>

[33mcommit d3fda51d11bc37fafc62a217a34d65287a54ec9a[m
Author: rameshthoomu <rameshbabu.thoomu@gmail.com>
Date:   Tue Dec 4 10:05:43 2018 -0500

    FAB-13155 Update multiarch script
    
    This patchset updates multiarch script to publish images
    with m.n version. This avoids changing the DockerImage tags
    in SDK docker-compose files for each release.
    
    Change-Id: I201be594fad302b51de2ccf9cce490100874d521
    Signed-off-by: rameshthoomu <rameshbabu.thoomu@gmail.com>

[33mcommit 081166aa3d66deef2b2953bd04c4321851986d35[m
Merge: 2c685b0 53b33d8
Author: Jonathan Levi (HACERA) <jonathan@hacera.com>
Date:   Mon Dec 10 13:17:47 2018 +0000

    Merge "[FAB-13005] Upgrade doc"

[33mcommit 1433a995df8c26bf8ec610af319dc140170784ad[m
Author: joe-alewine <Joe.Alewine@ibm.com>
Date:   Thu Dec 6 17:16:11 2018 -0500

    [FAB-13021] What's new in v1.4
    
    Rundown of new features and doc.
    
    Change-Id: I86bbb632c54b7dc197dd40423a04701e63698976
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit 2b6fcf46c5d03f92a93e2f4b155bd53d58031322[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Mon Dec 10 13:46:54 2018 +0200

    [FAB-13056] Onboarding: systemchannel from boot block
    
    This change set makes the system channel retrieval for onboarding be based
    on the content of the bootstrap block, and not in the orderer.yaml.
    
    Change-Id: I8009d079d080ac9ab855efa9dc322786605c0344
    Signed-off-by: yacovm <yacovm@il.ibm.com>

[33mcommit 713dc7a85e5e52680790edeb6e290c755279aed0[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Mon Dec 10 13:21:11 2018 +0200

    [FAB-13055] Don't close LedgerFactory in onboarding
    
    The ledgerFactory of the orderer should not be closed,
    since the golevel DB API clearly states that:
    
    "Other methods should not be called after the DB has been closed."
    
    Change-Id: I535c9ca7828292bc68f93f3be0cb110af98c956d
    Signed-off-by: yacovm <yacovm@il.ibm.com>

[33mcommit 8c2afba3f1ef58c569e5a719e807c26fcc49c817[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Mon Dec 10 13:09:47 2018 +0200

    [FAB-13054] Prevent int overflow in IsReplicationNeeded
    
    If the OSN is booted without a ledger, the height is 0.
    When deciding whether replication is needed, we comapre
    the numbers of the last committed block in the system channel.
    
    The height is always the last sequence + 1, therefore
    to get the last committed block number of the system channel
    we subtract 1 from the height, and when the height is 0-
    an integer overflow occurs.
    
    This change set fixes this, and adds a test case.
    
    Change-Id: Ia4718fb206670e50552f61168c73165e27e0fc68
    Signed-off-by: yacovm <yacovm@il.ibm.com>

[33mcommit 2c685b0a2a5b72730507f35c12e1cb06ae848927[m
Merge: fc25529 371dd0e
Author: Artem Barger <bartem@il.ibm.com>
Date:   Sun Dec 9 23:31:28 2018 +0000

    Merge "[FAB-13025] Logging, metrics, health doc"

[33mcommit 53b33d80e02b802ada2b282b958b9350df644978[m
Author: joe-alewine <Joe.Alewine@ibm.com>
Date:   Mon Dec 3 09:47:54 2018 -0500

    [FAB-13005] Upgrade doc
    
    Doc flow for upgrading to v1.4.
    
    Code/scripts available here:
    https://gerrit.hyperledger.org/r/c/27946/
    
    Change-Id: I3298896e8b35b82f5b04d8a3846c89a07fd807e1
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit fc2552930c6e165b16e12864b4ae867692e56220[m
Merge: 6f8d4ca 8f8d8dc
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Sun Dec 9 19:55:04 2018 +0000

    Merge "[FAB-13089] fetch attachments in CouchDB range queries"

[33mcommit 371dd0e0264c32b90ab70641b0006a824347f355[m
Author: joe-alewine <Joe.Alewine@ibm.com>
Date:   Fri Dec 7 16:59:27 2018 -0500

    [FAB-13025] Logging, metrics, health doc
    
    Main doc for conceptual info and setup.
    Metrics doc for autogenerated text.
    
    Change-Id: I89de4531b6e5be776d67f78d1739680550476472
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit 8f8d8dc899f42d78de760faa8a87b19f5a60fe43[m
Author: ChanderG <mail@chandergovind.org>
Date:   Sat Dec 8 16:20:35 2018 +0530

    [FAB-13089] fetch attachments in CouchDB range queries
    
    Using, attachments=true, fetch all attachments in a CouchDB range
    queries instead of first obtaining metadata and fetching each attachment
    in a separate call.
    
    Change-Id: Ibe51e621d9c49116ebd9fbd866f939eecff326a4
    Signed-off-by: ChanderG <mail@chandergovind.org>

[33mcommit 6f8d4cab7f8b3edb1beb29cae9af33d348964019[m
Merge: 8f5be7b d4e2016
Author: Manish Sethi <manish.sethi@gmail.com>
Date:   Sun Dec 9 03:12:05 2018 +0000

    Merge "[FAB-13161] Tips and tricks for couchdb"

[33mcommit d4e201634a8c04d9c97109f58d489b9e513f79ea[m
Author: joe-alewine <Joe.Alewine@ibm.com>
Date:   Tue Dec 4 16:58:23 2018 -0500

    [FAB-13161] Tips and tricks for couchdb
    
    Change-Id: I16a1cbff60637f1f610463a84f04514f614ac94c
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit 8f5be7b4dab9db739c805bc89bf2daa57ba502e1[m
Merge: d2511fe 09fe4c3
Author: Yacov Manevich <yacovm@il.ibm.com>
Date:   Sat Dec 8 19:21:19 2018 +0000

    Merge "FAB-13193 - Added help message to couch db metric"

[33mcommit d2511fea8bcb5f1521c752d15fe3da3fdf1bcadc[m
Merge: 2630e86 4b75195
Author: Yacov Manevich <yacovm@il.ibm.com>
Date:   Sat Dec 8 19:20:25 2018 +0000

    Merge "[FAB-13189] set content-type for logspec response"

[33mcommit 09fe4c3f789dc33f466e4ad2572cf913c03e962f[m
Author: Saad Karim <skarim@us.ibm.com>
Date:   Sat Dec 8 12:18:59 2018 -0500

    FAB-13193 - Added help message to couch db metric
    
    Added help message to couch db metric
    
    Change-Id: Icc24f39dcd8860ebe3872130a4318f848377dcc8
    Signed-off-by: Saad Karim <skarim@us.ibm.com>

[33mcommit 2630e86184dadd1bb6177b974f9d93cd111d7474[m
Merge: 7802c0e 2eff15f
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Sat Dec 8 17:17:42 2018 +0000

    Merge "Update sample configurations with operations"

[33mcommit 7802c0e9bce3bbb037cb6dd87112c25dbe4d8737[m
Merge: f4dad27 4f5fe21
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Sat Dec 8 17:08:29 2018 +0000

    Merge "[FAB-12947] More edits to gossip doc"

[33mcommit f4dad271ff666919adccf5f1033bbc568615a2e4[m
Merge: 9521e36 5802053
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Sat Dec 8 03:12:28 2018 +0000

    Merge "collACL: e2e test"

[33mcommit 4b751954f679b7ef5a30cd41e0b9aebc89e19732[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Fri Dec 7 18:45:12 2018 -0500

    [FAB-13189] set content-type for logspec response
    
    Change-Id: I4538e9c2f2f0f16e0a0362aa71df8496fe4228fd
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit 2eff15ffe262eae2052e47e67bfc0c1e7c6a0e82[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Fri Nov 30 11:36:13 2018 -0500

    Update sample configurations with operations
    
    FAB-13084 #done
    
    Change-Id: I808e7566227a59fa55a408c8b1e5e618f740218d
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit 9521e3681754d534d98c545f56b1eb1b69216500[m
Author: muralisr <srinivasan.muralidharan99@gmail.com>
Date:   Thu Dec 6 13:12:03 2018 -0500

    cleanup system chaincode comments in core.yaml
    
    This FAB is a minimal fix to remove the suggestion in core.yaml
    that users need to work with *core/scc/importsysccs.go* for creating
    system chaincode statically built into fabric.
    
    FAB-13185 #done
    Change-Id: I7a829551ab5ece7699bb4a217461ba1ba0d51a11
    Signed-off-by: muralisr <srinivasan.muralidharan99@gmail.com>

[33mcommit f9b4b11db4855ee5ab0acf38fdcbdd5375c1acc0[m
Merge: fe873f9 6b116aa
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Fri Dec 7 11:59:51 2018 +0000

    Merge "FAB-12896 Developing Apps: Gateway topic"

[33mcommit fe873f90e4c5a32a557c37b7b2982e7fcecf3dd4[m
Merge: 5fef8bd 684135e
Author: Manish Sethi <manish.sethi@gmail.com>
Date:   Fri Dec 7 05:18:23 2018 +0000

    Merge "FAB-13186 Fix couchdb version cache data race"

[33mcommit 684135efa9a311cfc74698309635c0ca4465a0b3[m
Author: manish <manish.sethi@gmail.com>
Date:   Thu Dec 6 19:22:51 2018 -0500

    FAB-13186 Fix couchdb version cache data race
    
    In one of the recent CRs
    https://gerrit.hyperledger.org/r/#/c/27580/11/core/ledger/kvledger/txmgmt/statedb/statecouchdb/statecouchdb.go
    a change in couchdb version cache loading code causes
    datarace between simulation and commit leading to
    concurrent modification of the underlying map
    
    This cr reverts this change
    
    Change-Id: I49fcad205d90b66786c2c0c48111af15ebb159e5
    Signed-off-by: manish <manish.sethi@gmail.com>

[33mcommit 5fef8bdae3bc590251ca94cbd1fb20bcda01f0e7[m
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Thu Dec 6 16:29:02 2018 -0500

    [FAB-11599] Private data doc edits
    
    Edits for private data reconciliation and
    access control.
    
    Change-Id: Ib094756644aebf01bb507c7ed057db3a8a13d72a
    Signed-off-by: David Enyeart <enyeart@us.ibm.com>
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit 4f5fe213428b7cb338bac1aa59bbc6b46a15ad71[m
Author: joe-alewine <Joe.Alewine@ibm.com>
Date:   Mon Dec 3 15:10:07 2018 -0500

    [FAB-12947] More edits to gossip doc
    
    Addressing additional comments from
    Dave E on this doc.
    
    Change-Id: I5180b20c076d563c5166cbdbd93a88731add45f5
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit fa9c5ee9c86cb9f297cff3bfed1437763449257c[m
Merge: 47868b8 2e4289e
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Thu Dec 6 20:21:56 2018 +0000

    Merge "[FAB-11599] private data remain doc"

[33mcommit 47868b85abedaeda0bd6bf0d40a624d7908e3d0d[m
Merge: 5b2b561 28fd96e
Author: Manish Sethi <manish.sethi@gmail.com>
Date:   Thu Dec 6 16:36:10 2018 +0000

    Merge "[FAB-13151] Fill in Length fields in attachments"

[33mcommit 2e4289e4d639bb71a575b63d1d170f3d768dd6c8[m
Author: joe-alewine <Joe.Alewine@ibm.com>
Date:   Wed Dec 5 12:06:33 2018 -0500

    [FAB-11599] private data remain doc
    
    Documenting reconciliation and
    access control in private data.
    
    Also did slight reorder on private
    data concept doc.
    
    Change-Id: I84c2b2d3f358d8f5568c7efa9f0ec61cc80c1e18
    Signed-off-by: joe-alewine <Joe.Alewine@ibm.com>

[33mcommit 5b2b5616d49538ace14aa9e08a482b9089c78df8[m
Author: Angelo De Caro <adc@zurich.ibm.com>
Date:   Wed Dec 5 12:40:09 2018 +0100

    [FAB-13139] Discard Idemix Empty Digest
    
    This change-set does the following:
    - remove the need for idemix empty digest
    
    Change-Id: I3d84284d290a735fae7ca71b464881b32c72294a
    Signed-off-by: Angelo De Caro <adc@zurich.ibm.com>

[33mcommit b62cf0ddb8d29198a153983d0f6b33838157a770[m
Merge: fb4de2f aabd259
Author: Artem Barger <bartem@il.ibm.com>
Date:   Wed Dec 5 22:57:52 2018 +0000

    Merge "[FAB-11746] Return cached alive message in gossip"

[33mcommit fb4de2fa4b3bc8112384353ed0429d88809c3423[m
Merge: a75a53d 5ef4e71
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 5 18:11:33 2018 +0000

    Merge "FAB-12496 Add disclaimer on Kafka/ZK security"

[33mcommit a75a53de4903727a1a2592358f681bbd86da0cd4[m
Merge: 55a5194 a578352
Author: Yacov Manevich <yacovm@il.ibm.com>
Date:   Wed Dec 5 16:24:54 2018 +0000

    Merge "FAB-13152 fix make protos"

[33mcommit 55a5194a9fdcf474a425854cae2af1f0dd509f89[m
Merge: 6a827fd 90f014e
Author: Artem Barger <bartem@il.ibm.com>
Date:   Wed Dec 5 16:16:09 2018 +0000

    Merge "[FAB-13164] Revert Fix pkcs11 UT Failures"

[33mcommit 6a827fd9c2206dff32d1aa6aae1509b38ba9ebe7[m
Merge: 5d3824e 300e509
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Dec 5 15:48:19 2018 +0000

    Merge "[FAB-12926] Validate hash chain when listing blocks"

[33mcommit 5d3824ecf73abba2b79290b053fff5ddb10f02e3[m
Merge: 2079f45 8939eb5
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Dec 5 15:07:42 2018 +0000

    Merge "acquire txSim only once during endorsement"

[33mcommit 90f014e578e045dfb05f5721ecca8d77ebf645e8[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Dec 4 15:41:26 2018 -0500

    [FAB-13164] Revert Fix pkcs11 UT Failures
    
    This reverts commit af5e9f0192be0013e0202c463328e13b9886a7e0.
    
    Change-Id: I8ef54da8f2040450fc82e44fd6c5089cc109a9f5
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit 2079f45689a4e29f36a31eaf01fe285271a15906[m
Merge: 3ed3426 f3d2130
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Dec 5 14:47:31 2018 +0000

    Merge "[FAB-12973] remove logging init from package init"

[33mcommit 3ed342680e64fb754b5be061f792797bc48b4101[m
Merge: d7d6313 386ea3d
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Dec 5 14:47:27 2018 +0000

    Merge "[FAB-13163] propagate caller environ to orderer"

[33mcommit d7d631380cfb69aa49a147707c05ed90869c0df6[m
Merge: 55d63d9 c85e7f9
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Dec 5 14:47:23 2018 +0000

    Merge "[FAB-12973] honor FABRIC_LOGGING_SPEC in library"

[33mcommit 55d63d90044a4274a0c15b3562c65f046d1787a5[m
Merge: 9034200 9dea224
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Dec 5 14:47:18 2018 +0000

    Merge "[FAB-13086] longer wait for term in healh check"

[33mcommit 903420046906ed92b332ac951d942adec822a99a[m
Merge: 98ab0fa 18551c4
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Wed Dec 5 14:47:03 2018 +0000

    Merge "[FAB-13025] Move namer from statsd to internal"

[33mcommit 98ab0fa927e97d3ba836528bc388b55785e187ca[m
Merge: e5e9e83 e0c729e
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 5 14:14:59 2018 +0000

    Merge "FAB-13158 Address data races in etcd/raft tests"

[33mcommit e5e9e8382f24a2982710ccf4dced2d338391d0af[m
Merge: 8d14cfc 3e7b2ef
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 5 13:26:46 2018 +0000

    Merge "pvtData APIs are not allowed in chaincode Init()"

[33mcommit a5783524e34a7ebe6d3e01698c7c945c6ae159cc[m
Author: Yoav Tock <tock@il.ibm.com>
Date:   Tue Dec 4 16:34:35 2018 +0200

    FAB-13152 fix make protos
    
    - update go_package path
    - update go:generate parameters
    - run 'make protos'
    
    Change-Id: Ie7c2825d274c4bcc12cf653d95de5b4fdaef2ade
    Signed-off-by: Yoav Tock <tock@il.ibm.com>

[33mcommit aabd259e09f3c2ab630d32db16cef2c32801e3ae[m
Author: Inbar Badian <inbar.badian@ibm.com>
Date:   Wed Dec 5 14:26:53 2018 +0200

    [FAB-11746] Return cached alive message in gossip
    
    Adding a cached field to the struct gossipDiscoveryImpl
    to avoid creating a SignedMessage every time
    
    Change-Id: I1e043ad5e3d929e9b08c5846e3d31bda55856540
    Signed-off-by: Inbar Badian <inbar.badian@ibm.com>

[33mcommit 8d14cfcdf8303dfc6a44164ead8f88f4ce08cadf[m
Merge: ece1498 acd9aaa
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 5 12:40:46 2018 +0000

    Merge "[FAB-13139] Fix Nym Public key Marshalling"

[33mcommit ece14986c1c3e545c7e848169d75649e5bea3ca6[m
Merge: d970db1 2ac523f
Author: Gari Singh <gari.r.singh@gmail.com>
Date:   Wed Dec 5 11:48:16 2018 +0000

    Merge "[FAB-13129] 1.4 release notes for operations"

[33mcommit acd9aaa6763f6afd16cba2815d5cf22b0698f095[m
Author: Angelo De Caro <adc@zurich.ibm.com>
Date:   Tue Dec 4 18:06:46 2018 +0100

    [FAB-13139] Fix Nym Public key Marshalling
    
    An Idemix-MSP serialised Identity contains the nym public key
    as two seperated fields for the X and Y coordinate.
    This change-set restores this by changing the way
    the BCCSP nym key is serialised.
    
    Change-Id: I6726335fdd02c237c8a75dd83886759968ea213a
    Signed-off-by: Angelo De Caro <adc@zurich.ibm.com>

[33mcommit d970db124999abea33c6a6f2cabfb91ae4db8636[m
Merge: 151f76e 42689db
Author: Alessandro Sorniotti <ale.linux@sopit.net>
Date:   Wed Dec 5 10:50:17 2018 +0000

    Merge "Fix handlers/validation UT failure on ppc64le"

[33mcommit 42689db8bc752231117fd4ae44315c44cbaf16fc[m
Author: Will Lahti <wtlahti@us.ibm.com>
Date:   Mon Dec 3 14:11:46 2018 -0500

    Fix handlers/validation UT failure on ppc64le
    
    FAB-12976 #done
    
    Change-Id: I0afcdb62579b258a0e5fc983745e7d937c77f843
    Signed-off-by: Will Lahti <wtlahti@us.ibm.com>

[33mcommit 151f76e4448916112e2733acd822260e054a58cd[m
Author: Angelo De Caro <adc@zurich.ibm.com>
Date:   Wed Dec 5 08:22:28 2018 +0100

    [FAB-13167] IdemixBridgeTest Robustification
    
    This change-set updates the tests to always
    correctly invalidate a credential in a bad path test.
    
    Change-Id: I3a0eda33bf1158d7a1cf8a3ddd7dc4c4af5548dd
    Signed-off-by: Angelo De Caro <adc@zurich.ibm.com>

[33mcommit 8939eb5214e355a4ce61a34892a251eaeb576d7a[m
Author: Senthil Nathan N <cendhu@gmail.com>
Date:   Wed Dec 5 00:02:40 2018 +0530

    acquire txSim only once during endorsement
    
    FAB-13026 #done
    
    Change-Id: I088b29b35a417ea9d23f748bf973ad37d215ea7d
    Signed-off-by: senthil <cendhu@gmail.com>

[33mcommit 3e7b2efd1a48b0996a46586dfaaac93fb8f3f754[m
Author: Senthil Nathan N <cendhu@gmail.com>
Date:   Tue Dec 4 20:48:27 2018 +0530

    pvtData APIs are not allowed in chaincode Init()
    
    private data APIs are not allowed in Init() since the
    private data collection configuration is not yet committed
    on the channel.
    
    This CR provide a better error message indicating that
    pvtdata APIs are  not allowed in chaincode Init().
    
    FAB-13050 #done
    
    Change-Id: I6e6a8abdf1f35f56307be90d8b6b2830c8f35755
    Signed-off-by: senthil <cendhu@gmail.com>

[33mcommit f63c95d3f7d83e64505c6d692adb23a651764e19[m
Merge: 385f437 16aecd0
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Wed Dec 5 04:32:18 2018 +0000

    Merge "Fix: Filter couchdb internal docs from query results"

[33mcommit 2ac523fa47539c7c2a8d092103cef16270f49c9f[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Dec 4 21:13:40 2018 -0500

    [FAB-13129] 1.4 release notes for operations
    
    Change-Id: I5dd53e4bb75ae4311156c4a729b3fceacad3efb4
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit e0c729ebba08580f8b4535ab1021601595c646ce[m
Author: Kostas Christidis <kostas@christidis.io>
Date:   Tue Dec 4 21:06:15 2018 -0500

    FAB-13158 Address data races in etcd/raft tests
    
    Change-Id: Ie918792893f3353652cc49f27622d02dd871e689
    Signed-off-by: Kostas Christidis <kostas@christidis.io>

[33mcommit 300e50916c692409ae269f785c759b2e6cce2c51[m
Author: yacovm <yacovm@il.ibm.com>
Date:   Tue Nov 20 13:22:28 2018 +0200

    [FAB-12926] Validate hash chain when listing blocks
    
    This change set adds validation of the hash chain to the ChainInspector.
    
    We need to validate the hash chain (backward hash pointers) when iterating over
    the system channel when inspecting the channels that exist (which is
    what the ChainInspector does).
    
    Even though the block puller that fetches the channels validates the hash chain in batches,
    it doesn't verify the hash pointers between the batches themselves.
    
    Since the onboarding puller (unlike the etcdraft one) is equipped with a no-op
    signature validator and relies on the validity of the bootstrap block
    in conjunction with the integrity of the pulled hash chain,
    it doesn't verify the signature of the last block of a pulled batch,
    thus we need to add this check into the code of the channel lister.
    
    Change-Id: I8320d093f8d6f5a81291e6377536b0ebdcc83c48
    Signed-off-by: yacovm <yacovm@il.ibm.com>

[33mcommit 385f4375659cb4a2c4af3631a3af889d335e17af[m
Merge: da77011 fd0263c
Author: Gari Singh <gari.r.singh@gmail.com>
Date:   Tue Dec 4 21:22:17 2018 +0000

    Merge "[FAB-13146] Modify the return value of NewMCS"

[33mcommit 5ef4e71efc5024d3b3c9d10bb43edf9bb7cb5d6d[m
Author: Kostas Christidis <kostas@christidis.io>
Date:   Tue Dec 4 16:21:05 2018 -0500

    FAB-12496 Add disclaimer on Kafka/ZK security
    
    Change-Id: I360c2a329d950a317222a2c139de29a1086ceccc
    Signed-off-by: Kostas Christidis <kostas@christidis.io>

[33mcommit f3d21304305835fadddf69fea4f462f8c44b99ab[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Dec 4 15:52:10 2018 -0500

    [FAB-12973] remove logging init from package init
    
    A package initializer buried in the configtxgen/localconfig package
    modifies the logging spec in the orderer at runtime. Remove it.
    
    Change-Id: I11bc289769b7069b0cad384b50ec8588e7f9a714
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit 386ea3d5d606c6b74e3eb1e2b21240a1ed194c5d[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Dec 4 15:02:52 2018 -0500

    [FAB-13163] propagate caller environ to orderer
    
    The environment should be propagated from the integration test framework
    to the processes that are spawned from it.
    
    Change-Id: Ib2c2bc3f0063f196af34d9e221db4a5d588b0793
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit c85e7f9af05d764f0b41fabd91377fb4da10bf6d[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Dec 4 14:59:56 2018 -0500

    [FAB-12973] honor FABRIC_LOGGING_SPEC in library
    
    When FABRIC_LOGGING_SPEC is set, the flogging library should use that
    value to initialize the logging library. When it is not set, the default
    logging spec of INFO should be used.
    
    Change-Id: I9c3e70f96b7da229fde7074969112021376307c8
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit 9dea22473c0c5a3955b30e8a24e25fa4fc05247e[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Dec 4 16:06:32 2018 -0500

    [FAB-13086] longer wait for term in healh check
    
    Also cleanup the dangling network and test directory.
    
    Change-Id: I145e4d97ed5bb0c7513c71ad8e0b7d6525f01128
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>

[33mcommit da77011232cfdcb16067956e01c4319e03769ee5[m
Merge: 5ab30b8 0cfb77b
Author: Gari Singh <gari.r.singh@gmail.com>
Date:   Tue Dec 4 20:58:15 2018 +0000

    Merge "FAB-12288 reduce log level extracting pvt data col."

[33mcommit fd0263cba36d59a12a22eddab5e060908c1c4396[m
Author: 乔伦 徐 <jamesxql@gmail.com>
Date:   Tue Dec 4 15:52:57 2018 +0800

    [FAB-13146] Modify the return value of NewMCS
    
    Change-Id: I64d1ba1bb3d2e917fdaaca61b34012c7febb342a
    Signed-off-by: 乔伦 徐 <jamesxql@gmail.com>

[33mcommit 5ab30b8d5336249ad12d84bb5703b3dcf83eaa39[m
Merge: 6c3895c de67216
Author: David Enyeart <enyeart@us.ibm.com>
Date:   Tue Dec 4 18:40:46 2018 +0000

    Merge "recon: rename CommitPvtData() ledger API"

[33mcommit 6c3895c4eec36be9813668fdd86e85b483b4802c[m
Merge: bfe01e4 c820bcb
Author: Kostas Christidis <kostas@gmail.com>
Date:   Tue Dec 4 17:57:07 2018 +0000

    Merge "[FAB-12708] 1/4 remove current orderer benchmark tests"

[33mcommit bfe01e4d1561ba20d8e6e2c7bc93302db9900044[m
Merge: 76b7d9c 95e4cde
Author: Yacov Manevich <yacovm@il.ibm.com>
Date:   Tue Dec 4 16:52:14 2018 +0000

    Merge "[FAB-13135,FAB-13136] Idemix/Fabric-CA Integration"

[33mcommit 76b7d9c9c2656b878291d37d79bcbc91912a1f31[m
Merge: 239155b 9d87d37
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Tue Dec 4 15:48:41 2018 +0000

    Merge "[FAB-13149] Disable etcdraft for v1.4"

[33mcommit 18551c438a939a93599216f7b1da841baf0da581[m
Author: Matthew Sykes <sykesmat@us.ibm.com>
Date:   Mon Dec 3 15:37:21 2018 -0500

    [FAB-13025] Move namer from statsd to internal
    
    Namer will be used as part of the metrics/gendoc tooling.
    
    Change-Id: I324ce50dd3df45afee6cb996798feec34f0cfc84
    Signed-off-by: Matthew Sykes <sykesmat@us.ibm.com>
