# #4090实现城市级3D重建##新架构挑战三维重建上限#用4张4090，就能搞定大规模城市三维重建？CityGS-X真做到了。这套由上海AI Lab与西工大提出的新架构，叫CityGS-...

**URL**: https://weibo.com/6105753431/PmWDK3fjt

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%234090%E5%AE%9E%E7%8E%B0%E5%9F%8E%E5%B8%82%E7%BA%A73D%E9%87%8D%E5%BB%BA%23&amp;extparam=%234090%E5%AE%9E%E7%8E%B0%E5%9F%8E%E5%B8%82%E7%BA%A73D%E9%87%8D%E5%BB%BA%23" data-hide=""><span class="surl-text">#4090实现城市级3D重建#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%96%B0%E6%9E%B6%E6%9E%84%E6%8C%91%E6%88%98%E4%B8%89%E7%BB%B4%E9%87%8D%E5%BB%BA%E4%B8%8A%E9%99%90%23&amp;extparam=%23%E6%96%B0%E6%9E%B6%E6%9E%84%E6%8C%91%E6%88%98%E4%B8%89%E7%BB%B4%E9%87%8D%E5%BB%BA%E4%B8%8A%E9%99%90%23" data-hide=""><span class="surl-text">#新架构挑战三维重建上限#</span></a><br><br>用4张4090，就能搞定大规模城市三维重建？CityGS-X真做到了。<br><br>这套由上海AI Lab与西工大提出的新架构，叫CityGS-X，核心是并行化混合分层三维表征（PH²-3D）。它绕开传统的“分块+合并”套路，靠批处理+多任务并行渲染，让训练更快、显存更省、精度更高。<br><br>几个亮点直接上干货：<br><br>- 训练速度翻倍，用的卡还更少。<br>    <br>- 渲染RGB图和还原几何精度，都能打当前SOTA。<br>    <br>- 即使是超5000张图的大场景，只要4卡4090就能跑。<br><br>模型设计上，他们用K层分层体素建模，每层都能并行处理，高效又精细。动态体素分配+高斯属性共享，显存控制做得很到位。<br><br>在渲染方面，模型按图块分配任务，支持高斯属性迁移，还用了法线和深度图协同优化。RGB训练、深度正则化和几何训练，都是批处理级别的，视角过拟合不再怕。 <a href="https://weibo.com/ttarticle/p/show?id=2309405154686322999548" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">4090玩转大场景几何重建，RGB渲染和几何精度达SOTA｜上海AI Lab＆西工</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0e8nk54wcj30d807ggm0.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

上海AI Lab与西工大联合提出CityGS-X架构，通过并行化混合分层三维表征（PH²-3D）实现高效城市级3D重建。该技术摒弃传统分块合并方法，采用批处理+多任务并行渲染，仅需4张RTX 4090即可处理5000+图像的大场景，训练速度翻倍且显存占用更低。其K层分层体素建模支持并行处理，结合动态体素分配与高斯属性共享，在RGB渲染和几何精度上均达到SOTA水平。通过图块任务分配、高斯属性迁移及法线/深度图协同优化，有效避免了视角过拟合问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-14T02:31:21Z
- **目录日期**: 2025-04-14
