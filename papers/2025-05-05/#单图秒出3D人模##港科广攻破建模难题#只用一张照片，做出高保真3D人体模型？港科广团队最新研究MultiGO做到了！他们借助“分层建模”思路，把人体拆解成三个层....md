# #单图秒出3D人模##港科广攻破建模难题#只用一张照片，做出高保真3D人体模型？港科广团队最新研究MultiGO做到了！他们借助“分层建模”思路，把人体拆解成三个层...

**URL**: https://weibo.com/6105753431/Pqp25l1rd

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%95%E5%9B%BE%E7%A7%92%E5%87%BA3D%E4%BA%BA%E6%A8%A1%23&amp;extparam=%23%E5%8D%95%E5%9B%BE%E7%A7%92%E5%87%BA3D%E4%BA%BA%E6%A8%A1%23" data-hide=""><span class="surl-text">#单图秒出3D人模#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%AF%E7%A7%91%E5%B9%BF%E6%94%BB%E7%A0%B4%E5%BB%BA%E6%A8%A1%E9%9A%BE%E9%A2%98%23&amp;extparam=%23%E6%B8%AF%E7%A7%91%E5%B9%BF%E6%94%BB%E7%A0%B4%E5%BB%BA%E6%A8%A1%E9%9A%BE%E9%A2%98%23" data-hide=""><span class="surl-text">#港科广攻破建模难题#</span></a><br><br>只用一张照片，做出高保真3D人体模型？港科广团队最新研究MultiGO做到了！<br><br>他们借助“分层建模”思路，把人体拆解成三个层级：骨架、关节、衣物褶皱，像搭乐高一样从大块建轮廓，到小块补细节，再上纹理，层层细化。研究已入选CVPR 2025，代码也即将开源。<br><br>以前的方法多靠预训练模板，难搞细节、不准骨骼。MultiGO用三招解决：<br><br>- 骨架增强：把3D傅里叶特征投影到2D图像，借SMPL-X模板搞定骨架对齐。<br>- 关节扰动：训练时特意“搞乱”关节位置，让模型更抗误差。<br>- 皱纹去噪：用扩散模型思路，把粗糙当噪声处理，细节一点点“洗”出来。<br><br>三层模块层层递进，互相配合。骨架准了，关节稳了，皱纹细了，模型从整体到细节都更像人了。最终效果直接拉满，在CustomHuman和THuman3.0测试集中，精度指标全线领先。<br><br>不仅精度高，这项技术还有超多用武之地，如虚拟试衣、虚拟角色定制、影视特效等，详情看文章：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FYBfP4Mr1byrM2JtRbI5gZg" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a> <a href="https://weibo.com/ttarticle/p/show?id=2309405162927576711228" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">细节直逼亚毫米级！港科广分层建模突破3D人体生成｜CVPR 2025</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i14j5dp8wlj30bk06idgc.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

港科广团队提出MultiGO方法，通过分层建模（骨架、关节、衣物褶皱三层次）实现单张照片生成高保真3D人体模型。该方法采用骨架增强（3D傅里叶特征投影）、关节扰动训练和皱纹去噪扩散模型三项关键技术，解决了传统模板法细节缺失和骨骼不准的问题。在CustomHuman和THuman3.0测试中精度领先，适用于虚拟试衣、影视特效等场景。研究入选CVPR 2025，代码即将开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T05:02:44Z
- **目录日期**: 2025-05-05
