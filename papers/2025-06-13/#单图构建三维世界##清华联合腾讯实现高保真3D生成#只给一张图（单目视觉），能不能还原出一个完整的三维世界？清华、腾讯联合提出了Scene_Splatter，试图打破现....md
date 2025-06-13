# #单图构建三维世界##清华联合腾讯实现高保真3D生成#只给一张图（单目视觉），能不能还原出一个完整的三维世界？清华、腾讯联合提出了Scene Splatter，试图打破现...

**URL**: https://weibo.com/6105753431/PwlBSbGRt

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%95%E5%9B%BE%E6%9E%84%E5%BB%BA%E4%B8%89%E7%BB%B4%E4%B8%96%E7%95%8C%23&amp;extparam=%23%E5%8D%95%E5%9B%BE%E6%9E%84%E5%BB%BA%E4%B8%89%E7%BB%B4%E4%B8%96%E7%95%8C%23" data-hide=""><span class="surl-text">#单图构建三维世界#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%85%E5%8D%8E%E8%81%94%E5%90%88%E8%85%BE%E8%AE%AF%E5%AE%9E%E7%8E%B0%E9%AB%98%E4%BF%9D%E7%9C%9F3D%E7%94%9F%E6%88%90%23&amp;extparam=%23%E6%B8%85%E5%8D%8E%E8%81%94%E5%90%88%E8%85%BE%E8%AE%AF%E5%AE%9E%E7%8E%B0%E9%AB%98%E4%BF%9D%E7%9C%9F3D%E7%94%9F%E6%88%90%23" data-hide=""><span class="surl-text">#清华联合腾讯实现高保真3D生成#</span></a><br><br>只给一张图（单目视觉），能不能还原出一个完整的三维世界？<br><br>清华、腾讯联合提出了Scene Splatter，试图打破现有方法在三维一致性和细节恢复上的瓶颈。<br><br>现状是这样的：<br><br>- 主流三维生成模型（比如Hunyuan3D、Rodin-v1.5、Tripo-v2.5）大多只擅长单个物体的构建；<br>    <br>- 但一旦扩展到复杂场景，就会出现结构扭曲、几何缺失等问题，尤其在输入只有一张图的情况下更是“病态问题”。<br>    <br>Scene Splatter的关键突破在于：<br><br>1. 引入“动量引导”思想：<br>    - 类似优化算法中的“动量”机制，把上一次生成的细节引入当前步骤；<br>    - 首先用潜空间动量补细节，再用像素级动量补未知区域，实现高保真多视角视频生成。<br>        <br>2. 基于视频而非单图重建三维场景：<br>    - 通过视频扩散模型“合成”出多个视角，再用这些视角做三维恢复；<br>    - 大大提升了三维一致性和细节还原度。<br>        <br>3. 可泛化到任意相机轨迹：<br>    - 不管是环绕、拉远还是旋转视角，Scene Splatter都能生成连贯、稳定的三维视频。<br><br>团队还通过消融实验验证了动量机制的必要性：缺了动量，PSNR和SSIM指标都有显著下降，说明场景一致性和生成质量都会受影响。 <a href="https://weibo.com/ttarticle/p/show?id=2309405177095679770735" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">视频扩散模型新突破！清华腾讯联合实现高保真3D生成，告别多视图依赖</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2dpjplrxhj30ow0e0414.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

清华大学与腾讯联合提出Scene Splatter技术，突破单图构建三维世界的瓶颈。现有3D生成模型多局限于单个物体，复杂场景易出现结构扭曲。该技术引入"动量引导"机制，通过潜空间和像素级动量补全细节，并基于视频扩散模型合成多视角提升3D一致性。实验表明，动量机制对保持场景连贯性至关重要，支持任意相机轨迹生成稳定3D视频，显著提高了PSNR和SSIM指标。这一进展为单图高保真3D重建提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T09:03:57Z
- **目录日期**: 2025-06-13
