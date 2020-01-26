package com.documentqlite.base;

public interface Page {

    public byte[] getPage(int pageNum);

    public void writePage(int pageNum, byte[] date);

    public void deletePage(int pageNum);

    public long getMaximumPageNumber();

    public long createNewPage(byte[] data);
}
