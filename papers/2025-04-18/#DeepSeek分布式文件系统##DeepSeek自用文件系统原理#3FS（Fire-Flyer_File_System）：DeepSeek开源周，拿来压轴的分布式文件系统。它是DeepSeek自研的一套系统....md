# #DeepSeek分布式文件系统##DeepSeek自用文件系统原理#3FS（Fire-Flyer File System）：DeepSeek开源周，拿来压轴的分布式文件系统。它是DeepSeek自研的一套系统...

**URL**: https://weibo.com/6105753431/PnNMw29Tm

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23DeepSeek%E5%88%86%E5%B8%83%E5%BC%8F%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%23&amp;extparam=%23DeepSeek%E5%88%86%E5%B8%83%E5%BC%8F%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%23" data-hide=""><span class="surl-text">#DeepSeek分布式文件系统#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23DeepSeek%E8%87%AA%E7%94%A8%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%E5%8E%9F%E7%90%86%23&amp;extparam=%23DeepSeek%E8%87%AA%E7%94%A8%E6%96%87%E4%BB%B6%E7%B3%BB%E7%BB%9F%E5%8E%9F%E7%90%86%23" data-hide=""><span class="surl-text">#DeepSeek自用文件系统原理#</span></a><br><br>3FS（Fire-Flyer File System）：DeepSeek开源周，拿来压轴的分布式文件系统。<br><br>它是DeepSeek自研的一套系统，V3和R1训练时用的就是它。<br><br>该系统能够极大地发挥固态硬盘的带宽优势，在大规模集群环境下展现出强大的读写能力和容错性能。<br><br>在180个节点组成的集群中，3FS可实现高达6.6TB/s的聚合读取吞吐量。<br><br>而且3FS还对用户高度透明。就算一个文件存散在10台机器上，用户看到的依然只是一个/3fs/xxx.txt路径，读写方式和本地没区别。【图1】<br><br>在具体实现上，3FS中的节点分为四种角色：【图2】<br><br>- Client：用户接入的入口节点，负责与其他节点协调交互；<br>- Meta：负责管理文件的元数据，包括路径、权限、文件大小等；<br>- Mgmtd：集群的调度核心，掌握节点分布情况，并负责数据复制与节点注册；<br>- Storage：实际存储数据块的节点。<br><br>在这种架构下，3FS的核心目标是实现“强一致性”——即使文件被分布在多个节点，用户看到的路径和内容始终一致。<br><br>实现强一致性的关键在于引入CRAQ（Chain Replication with Apportioned Queries）协议。<br><br>CRAQ通过链式复制的方式，有序传递写操作，确保数据的一致性。其工作流程如下：【图3】<br><br>- 写入开始：用户发起写操作（如name=henry），数据首先写入链的头节点（Head）；<br>- 沿链传播：数据依次传递到中间副本节点（Replica）和尾节点（Tail）；<br>- 写入状态变化：<br>  - 数据在传输过程中被标记为Dirty（脏数据，不可读取）；<br>  - 一旦数据成功写入尾节点，便被标记为Clean（清洁数据，可读取）；<br>- 尾节点读取：为保证读取的是一致且已确认的数据，读操作通常只在尾节点进行。<br><br>CRAQ的好处是读性能好，但由于写入过程需经过整个链条，其写入延迟可能受到最慢节点的影响。因此，在读写都非常密集的场景中，可能面临性能瓶颈。<br><br>目前，该项目还在持续完善中，感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fdeepseek-ai%2F3FS" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br><br>扩展阅读：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmaknee.github.io%2Fblog%2F2025%2F3FS-Performance-Journal-1%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0kr9izvb9j31zv0dqtmh.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0kr9k5pa5j30mr0e5wjj.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0kr9lvmq0j30gf08u0v2.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

DeepSeek开源了自研的分布式文件系统3FS，该系统在180节点集群中可实现6.6TB/s的聚合读取吞吐量，专为发挥SSD带宽优势而设计。3FS采用四层架构（Client、Meta、Mgmtd、Storage），通过CRAQ协议实现强一致性：写入数据从Head节点经Replica传递至Tail节点，由Dirty状态转为Clean后才可读。虽然链式复制保证了数据一致性，但可能影响写入延迟。系统对用户透明，分布式存储的文件仍以统一路径呈现。该项目已用于DeepSeek的V3/R1模型训练，目前仍在持续优化中。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-18T10:03:58Z
- **目录日期**: 2025-04-18
